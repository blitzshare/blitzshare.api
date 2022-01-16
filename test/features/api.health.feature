Feature: Validate API health
  Scenario: Api is deployed and responsive on health check endpoint
  Given http GET health check request executed
  Then http response is OK