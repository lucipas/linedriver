
# COMPLETED FEATURES
DELETE /session/{session id}/actions                                     Release Actions
DELETE /session/{session id}/cookie                                      Delete All Cookies
DELETE /session/{session id}/cookie/{name}                               Delete Cookie
DELETE /session/{session id}                                             Delete Session
DELETE /session/{session id}/window                                      Close Window

GET /session/{session id}/alert/text                                     Get Alert Text
GET /session/{session id}/cookie                                         Get All Cookies
GET /session/{session id}/cookie/{name}                                  Get Named Cookie
GET /session/{session id}/element/active                                 Get Active Element
GET /session/{session id}/element/{element id}/attribute/{name}          Get Element Attribute
GET /session/{session id}/element/{element id}/computedlabel             Get Computed Label
GET /session/{session id}/element/{element id}/computedrole              Get Computed Role
GET /session/{session id}/element/{element id}/css/{property name}       Get Element CSS Value
GET /session/{session id}/element/{element id}/enabled                   Is Element Enabled
GET /session/{session id}/element/{element id}/name                      Get Element Tag Name
GET /session/{session id}/element/{element id}/property/{name}           Get Element Property
GET /session/{session id}/element/{element id}/rect                      Get Element Rect
GET /session/{session id}/element/{element id}/screenshot                Take Element Screenshot
GET /session/{session id}/element/{element id}/selected                  Is Element Selected
GET /session/{session id}/element/{element id}/shadow                    Get Element Shadow Root
GET /session/{session id}/element/{element id}/text                      Get Element Text
GET /session/{session id}/screenshot                                     Take Screenshot
GET /session/{session id}/source                                         Get Page Source
GET /session/{session id}/timeouts                                       Get Timeouts
GET /session/{session id}/title Get                                      Title
GET /session/{session id}/url                                            Get Current URL
GET /session/{session id}/window                                         Get Window Handle
GET /session/{session id}/window/handles                                 Get Window Handles
GET /session/{session id}/window/rect                                    Get Window Rect
GET /status                                                              Status

POST /session                                                            New Session
POST /session/{session id}/actions                                       Perform Actions
POST /session/{session id}/alert/accept                                  Accept Alert
POST /session/{session id}/alert/dismiss                                 Dismiss Alert
POST /session/{session id}/alert/text                                    Send Alert Text
POST /session/{session id}/back                                          Back
POST /session/{session id}/cookie                                        Add Cookie
POST /session/{session id}/element/{element id}/clear                    Element Clear
POST /session/{session id}/element/{element id}/click                    Element Click
POST /session/{session id}/element/{element id}/element                  Find Element From Element
POST /session/{session id}/element/{element id}/elements                 Find Elements From Element
POST /session/{session id}/element/{element id}/value                    Element Send Keys
POST /session/{session id}/element                                       Find Element
POST /session/{session id}/elements                                      Find Elements
POST /session/{session id}/execute/async                                 Execute Async Script
POST /session/{session id}/execute/sync                                  Execute Script
POST /session/{session id}/forward                                       Forward
POST /session/{session id}/frame/parent                                  Switch To Parent Frame
POST /session/{session id}/frame                                         Switch To Frame
POST /session/{session id}/print                                         Print
POST /session/{session id}/refresh                                       Refresh
POST /session/{session id}/shadow/{shadow id}/element                    Find Element From Shadow Root
POST /session/{session id}/shadow/{shadow id}/elements                   Find Elements From Shadow Root
POST /session/{session id}/timeouts                                      Set Timeouts
POST /session/{session id}/url                                           Navigate To
POST /session/{session id}/window/fullscreen                             Fullscreen Window
POST /session/{session id}/window/maximize                               Maximize Window
POST /session/{session id}/window/minimize                               Minimize Window
POST /session/{session id}/window/new                                    New Window
POST /session/{session id}/window/rect                                   Set Window Rect
POST /session/{session id}/window                                        Switch To Window
