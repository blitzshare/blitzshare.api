Feature: API peers connectivity
  Scenario: Peer can
  Given User registers via OTP
  When Another User obtains registred user information via OTP
  And User get bootstrap node config
  Then Connection between useres can be etablished
  And User can deregister OTP via obtained Token