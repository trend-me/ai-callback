Feature: Build AI prompts and forward the event to ai-prompt-builder queue

  Scenario: Successfully process a message from the queue
    Given a message with the following data is sent to 'ai-callback' queue:
    """
    {
    "prompt_road_map_config_execution_id":"c713deb9-efa2-4d5f-9675-abe0b7e0c0d4",
    "prompt_road_map_config_name":"TEST",
    "output_queue":"output-queue",
    "prompt_road_map_step":2,
    "model":"GEMINI",
    "metadata":{"any": { "thing":"test", "array":[1,2,3,4]} }
    }
    """
    Given the prompt road map API returns the following prompt road map for step '3' and prompt_road_map_config_name 'TEST':
    """
    {
    "prompt_road_map_config_name":"TEST",
    "response_validation_name":"TEST_RESPONSE",
    "metadata_validation_name":"TEST_METADATA",
    "question_template":"this is a <any.thing>. <any.array> <any.array[0]>",
    "step":3,
    "created_at":"2024-08-01T20:53:49.132Z",
    "updated_at":"2024-08-01T20:53:49.132Z"
    }
    """
    When the message is consumed by the ai-callback consumer
    Then the prompt_road_map is fetched from the prompt-road-map-api using the prompt_road_map_config_name 'TEST' and step '3'
    And the prompt_road_map_config_execution step_in_execution is updated to '3'
    And a message with the following data should be sent to 'ai-prompt-builder' queue:
    """
    {
    "prompt_road_map_config_execution_id":"c713deb9-efa2-4d5f-9675-abe0b7e0c0d4",
    "prompt_road_map_config_name":"TEST",
    "output_queue":"output-queue",
    "model":"GEMINI",
    "prompt_road_map_step":3,
    "metadata":{"any": { "thing":"test", "array":[1,2,3,4]} }
    }
    """
    And no message should be sent to the 'output-queue' queue
    And the application should not retry

  
  Scenario: Successfully sends a message to the output queue when the last road map step arrives
    Given a message with the following data is sent to 'ai-callback' queue:
    """
    {
    "prompt_road_map_config_execution_id":"c713deb9-efa2-4d5f-9675-abe0b7e0c0d4",
    "prompt_road_map_config_name":"TEST",
    "output_queue":"output-queue",
    "prompt_road_map_step":2,
    "model":"GEMINI",
    "metadata":{"any": { "thing":"test", "array":[1,2,3,4]} }
    }
    """
    Given the prompt road map API returns an statusCode '404'
    When the message is consumed by the ai-callback consumer
    Then the prompt_road_map is fetched from the prompt-road-map-api using the prompt_road_map_config_name 'TEST' and step '3'
    And no prompt_road_map_config_execution should be updated
    And a message with the following data should be sent to 'output-queue' queue:
    """
    {
    "prompt_road_map_config_execution_id":"c713deb9-efa2-4d5f-9675-abe0b7e0c0d4",
    "prompt_road_map_config_name":"TEST",
    "output_queue":"output-queue",
    "model":"GEMINI",
    "prompt_road_map_step":2,
    "metadata":{"any": { "thing":"test", "array":[1,2,3,4]} }
    }
    """
    And no message should be sent to the 'ai-prompt-builder' queue
    And the application should not retry


  Scenario: Successfully process an error and scheduling a retry
    Given a message with the following data is sent to 'ai-callback' queue:
    """
    {
    "prompt_road_map_config_execution_id":"c713deb9-efa2-4d5f-9675-abe0b7e0c0d4",
    "prompt_road_map_config_name":"TEST",
    "output_queue":"output-queue",
    "prompt_road_map_step":2,
    "model":"GEMINI",
    "metadata":{"any": { "thing":"test", "array":[1,2,3,4]} }
    }
    """
    Given the prompt road map API returns an statusCode '500'
    When the message is consumed by the ai-callback consumer
    Then the prompt_road_map is fetched from the prompt-road-map-api using the prompt_road_map_config_name 'TEST' and step '3'
    And no prompt_road_map_config_execution should be updated
    And no message should be sent to the 'ai-prompt-builder' queue
    And no message should be sent to the 'output-queue' queue
    And the application should retry

  Scenario: Successfully sends an error to output queue when max retry attempts is reached
    Given a message with the following data is sent to 'ai-callback' queue:
    """
    {
    "prompt_road_map_config_execution_id":"c713deb9-efa2-4d5f-9675-abe0b7e0c0d4",
    "prompt_road_map_config_name":"TEST",
    "output_queue":"output-queue",
    "prompt_road_map_step":2,
    "model":"GEMINI",
    "metadata":{"any": { "thing":"test", "array":[1,2,3,4]} }
    }
    """
    Given the prompt road map API returns an statusCode '500'
    Given max receive count is '-1'
    When the message is consumed by the ai-callback consumer
    Then the prompt_road_map is fetched from the prompt-road-map-api using the prompt_road_map_config_name 'TEST' and step '3'
    And no prompt_road_map_config_execution should be updated
    And no message should be sent to the 'ai-prompt-builder' queue
    And a message with the following data should be sent to 'output-queue' queue:
    """
    {
    "prompt_road_map_config_execution_id":"c713deb9-efa2-4d5f-9675-abe0b7e0c0d4",
    "prompt_road_map_config_name":"TEST",
    "output_queue":"output-queue",
    "model":"GEMINI",
    "prompt_road_map_step":2,
    "metadata":{"any": { "thing":"test", "array":[1,2,3,4]} },
    "error": {
      "message": ["response with statusCode: '500 Internal Server Error'"],
      "error_type":"Get Prompt Road Map Config Error",
      "abort": false,
      "notify": true 
      }
    }
    """
    And the application should not retry