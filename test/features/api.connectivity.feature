Feature: Peers connectivity
   Scenario: Users can obtain information inorder to connect to each other
    Given User registers via OTP
    When Another User obtains registred user information via OTP
    And User get bootstrap node config
    Then Connection between useres can be etablished
  
  Scenario: User can deregister OTP via obtained Token
    Given User registers via OTP
    Then User can deregister OTP via obtained Token