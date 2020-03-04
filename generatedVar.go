package gmb

var (
	dllGetHandleVerifier                                   uintptr
	dllJsAddRef                                            uintptr
	dllJsArg                                               uintptr
	dllJsArgCount                                          uintptr
	dllJsArgType                                           uintptr
	dllJsArrayBuffer                                       uintptr
	dllJsBindFunction                                      uintptr
	dllJsBindGetter                                        uintptr
	dllJsBindSetter                                        uintptr
	dllJsBoolean                                           uintptr
	dllJsCall                                              uintptr
	dllJsCallGlobal                                        uintptr
	dllJsDeleteObjectProp                                  uintptr
	dllJsDouble                                            uintptr
	dllJsEmptyArray                                        uintptr
	dllJsEmptyObject                                       uintptr
	dllJsEval                                              uintptr
	dllJsEvalExW                                           uintptr
	dllJsEvalW                                             uintptr
	dllJsFalse                                             uintptr
	dllJsFloat                                             uintptr
	dllJsFunction                                          uintptr
	dllJsGC                                                uintptr
	dllJsGet                                               uintptr
	dllJsGetArrayBuffer                                    uintptr
	dllJsGetAt                                             uintptr
	dllJsGetCallstack                                      uintptr
	dllJsGetData                                           uintptr
	dllJsGetGlobal                                         uintptr
	dllJsGetKeys                                           uintptr
	dllJsGetLastErrorIfException                           uintptr
	dllJsGetLength                                         uintptr
	dllJsGetWebView                                        uintptr
	dllJsGlobalObject                                      uintptr
	dllJsInt                                               uintptr
	dllJsIsArray                                           uintptr
	dllJsIsBoolean                                         uintptr
	dllJsIsFalse                                           uintptr
	dllJsIsFunction                                        uintptr
	dllJsIsJsValueValid                                    uintptr
	dllJsIsNull                                            uintptr
	dllJsIsNumber                                          uintptr
	dllJsIsObject                                          uintptr
	dllJsIsString                                          uintptr
	dllJsIsTrue                                            uintptr
	dllJsIsUndefined                                       uintptr
	dllJsIsValidExecState                                  uintptr
	dllJsNull                                              uintptr
	dllJsObject                                            uintptr
	dllJsReleaseRef                                        uintptr
	dllJsSet                                               uintptr
	dllJsSetAt                                             uintptr
	dllJsSetGlobal                                         uintptr
	dllJsSetLength                                         uintptr
	dllJsString                                            uintptr
	dllJsStringW                                           uintptr
	dllJsThrowException                                    uintptr
	dllJsToBoolean                                         uintptr
	dllJsToDouble                                          uintptr
	dllJsToFloat                                           uintptr
	dllJsToInt                                             uintptr
	dllJsToString                                          uintptr
	dllJsToStringW                                         uintptr
	dllJsToTempString                                      uintptr
	dllJsToTempStringW                                     uintptr
	dllJsToV8Value                                         uintptr
	dllJsTrue                                              uintptr
	dllJsTypeOf                                            uintptr
	dllJsUndefined                                         uintptr
	dllWkeAddDirtyArea                                     uintptr
	dllWkeAddNpapiPlugin                                   uintptr
	dllWkeAddPluginDirectory                               uintptr
	dllWkeAwaken                                           uintptr
	dllWkeCanGoBack                                        uintptr
	dllWkeCanGoForward                                     uintptr
	dllWkeClearCookie                                      uintptr
	dllWkeConfigure                                        uintptr
	dllWkeContentsHeight                                   uintptr
	dllWkeContentsWidth                                    uintptr
	dllWkeContextMenuEvent                                 uintptr
	dllWkeCookieEnabled                                    uintptr
	dllWkeCopy                                             uintptr
	dllWkeCreateMemBuf                                     uintptr
	dllWkeCreateString                                     uintptr
	dllWkeCreateStringW                                    uintptr
	dllWkeCreateStringWithoutNullTermination               uintptr
	dllWkeCreateWebCustomWindow                            uintptr
	dllWkeCreateWebView                                    uintptr
	dllWkeCreateWebWindow                                  uintptr
	dllWkeCut                                              uintptr
	dllWkeDelete                                           uintptr
	dllWkeDeleteString                                     uintptr
	dllWkeDestroyWebView                                   uintptr
	dllWkeDestroyWebWindow                                 uintptr
	dllWkeDragTargetDragEnter                              uintptr
	dllWkeDragTargetDragLeave                              uintptr
	dllWkeDragTargetDragOver                               uintptr
	dllWkeDragTargetDrop                                   uintptr
	dllWkeDragTargetEnd                                    uintptr
	dllWkeEditorCopy                                       uintptr
	dllWkeEditorCut                                        uintptr
	dllWkeEditorDelete                                     uintptr
	dllWkeEditorPaste                                      uintptr
	dllWkeEditorRedo                                       uintptr
	dllWkeEditorSelectAll                                  uintptr
	dllWkeEditorUnSelect                                   uintptr
	dllWkeEditorUndo                                       uintptr
	dllWkeEnableWindow                                     uintptr
	dllWkeFinalize                                         uintptr
	dllWkeFireContextMenuEvent                             uintptr
	dllWkeFireKeyDownEvent                                 uintptr
	dllWkeFireKeyPressEvent                                uintptr
	dllWkeFireKeyUpEvent                                   uintptr
	dllWkeFireMouseEvent                                   uintptr
	dllWkeFireMouseWheelEvent                              uintptr
	dllWkeFireWindowsMessage                               uintptr
	dllWkeFocus                                            uintptr
	dllWkeFreeMemBuf                                       uintptr
	dllWkeGC                                               uintptr
	dllWkeGetBlinkMainThreadIsolate                        uintptr
	dllWkeGetCaret                                         uintptr
	dllWkeGetCaretRect                                     uintptr
	dllWkeGetClientHandler                                 uintptr
	dllWkeGetContentAsMarkup                               uintptr
	dllWkeGetContentHeight                                 uintptr
	dllWkeGetContentWidth                                  uintptr
	dllWkeGetCookie                                        uintptr
	dllWkeGetCookieW                                       uintptr
	dllWkeGetCursorInfoType                                uintptr
	dllWkeGetDebugConfig                                   uintptr
	dllWkeGetDocumentCompleteURL                           uintptr
	dllWkeGetFrameUrl                                      uintptr
	dllWkeGetGlobalExecByFrame                             uintptr
	dllWkeGetHeight                                        uintptr
	dllWkeGetHostHWND                                      uintptr
	dllWkeGetMediaVolume                                   uintptr
	dllWkeGetName                                          uintptr
	dllWkeGetSource                                        uintptr
	dllWkeGetString                                        uintptr
	dllWkeGetStringLen                                     uintptr
	dllWkeGetStringW                                       uintptr
	dllWkeGetTempCallbackInfo                              uintptr
	dllWkeGetTitle                                         uintptr
	dllWkeGetTitleW                                        uintptr
	dllWkeGetURL                                           uintptr
	dllWkeGetUserAgent                                     uintptr
	dllWkeGetUserKeyValue                                  uintptr
	dllWkeGetVersion                                       uintptr
	dllWkeGetVersionString                                 uintptr
	dllWkeGetViewDC                                        uintptr
	dllWkeGetWebViewByNData                                uintptr
	dllWkeGetWebViewForCurrentContext                      uintptr
	dllWkeGetWebviewId                                     uintptr
	dllWkeGetWidth                                         uintptr
	dllWkeGetWindowHandle                                  uintptr
	dllWkeGetZoomFactor                                    uintptr
	dllWkeGlobalExec                                       uintptr
	dllWkeGoBack                                           uintptr
	dllWkeGoForward                                        uintptr
	dllWkeGoToIndex                                        uintptr
	dllWkeGoToOffset                                       uintptr
	dllWkeHeight                                           uintptr
	dllWkeInit                                             uintptr
	dllWkeInitialize                                       uintptr
	dllWkeInitializeEx                                     uintptr
	dllWkeInsertCSSByFrame                                 uintptr
	dllWkeIsAwake                                          uintptr
	dllWkeIsCookieEnabled                                  uintptr
	dllWkeIsDirty                                          uintptr
	dllWkeIsDocumentReady                                  uintptr
	dllWkeIsInitialize                                     uintptr
	dllWkeIsLoadComplete                                   uintptr
	dllWkeIsLoadFailed                                     uintptr
	dllWkeIsLoaded                                         uintptr
	dllWkeIsLoading                                        uintptr
	dllWkeIsLoadingCompleted                               uintptr
	dllWkeIsLoadingFailed                                  uintptr
	dllWkeIsLoadingSucceeded                               uintptr
	dllWkeIsMainFrame                                      uintptr
	dllWkeIsProcessingUserGesture                          uintptr
	dllWkeIsTransparent                                    uintptr
	dllWkeIsWebRemoteFrame                                 uintptr
	dllWkeIsWebviewAlive                                   uintptr
	dllWkeIsWebviewValid                                   uintptr
	dllWkeJsBindFunction                                   uintptr
	dllWkeJsBindGetter                                     uintptr
	dllWkeJsBindSetter                                     uintptr
	dllWkeKeyDown                                          uintptr
	dllWkeKeyPress                                         uintptr
	dllWkeKeyUp                                            uintptr
	dllWkeKillFocus                                        uintptr
	dllWkeLayoutIfNeeded                                   uintptr
	dllWkeLoadFile                                         uintptr
	dllWkeLoadFileW                                        uintptr
	dllWkeLoadHTML                                         uintptr
	dllWkeLoadHTMLW                                        uintptr
	dllWkeLoadHtmlWithBaseUrl                              uintptr
	dllWkeLoadURL                                          uintptr
	dllWkeLoadURLW                                         uintptr
	dllWkeLoadW                                            uintptr
	dllWkeMediaVolume                                      uintptr
	dllWkeMouseEvent                                       uintptr
	dllWkeMouseWheel                                       uintptr
	dllWkeMoveToCenter                                     uintptr
	dllWkeMoveWindow                                       uintptr
	dllWkeNetAddHTTPHeaderFieldToUrlRequest                uintptr
	dllWkeNetCancelRequest                                 uintptr
	dllWkeNetCancelWebUrlRequest                           uintptr
	dllWkeNetChangeRequestUrl                              uintptr
	dllWkeNetContinueJob                                   uintptr
	dllWkeNetCopyWebUrlRequest                             uintptr
	dllWkeNetCreatePostBodyElement                         uintptr
	dllWkeNetCreatePostBodyElements                        uintptr
	dllWkeNetCreateWebUrlRequest                           uintptr
	dllWkeNetCreateWebUrlRequest2                          uintptr
	dllWkeNetDeleteBlinkWebURLRequestPtr                   uintptr
	dllWkeNetFreePostBodyElement                           uintptr
	dllWkeNetFreePostBodyElements                          uintptr
	dllWkeNetGetExpectedContentLength                      uintptr
	dllWkeNetGetFavicon                                    uintptr
	dllWkeNetGetHTTPHeaderField                            uintptr
	dllWkeNetGetHTTPHeaderFieldFromResponse                uintptr
	dllWkeNetGetHttpStatusCode                             uintptr
	dllWkeNetGetMIMEType                                   uintptr
	dllWkeNetGetPostBody                                   uintptr
	dllWkeNetGetRawHttpHead                                uintptr
	dllWkeNetGetRawResponseHead                            uintptr
	dllWkeNetGetRequestMethod                              uintptr
	dllWkeNetGetResponseUrl                                uintptr
	dllWkeNetGetUrlByJob                                   uintptr
	dllWkeNetHoldJobToAsynCommit                           uintptr
	dllWkeNetHookRequest                                   uintptr
	dllWkeNetOnResponse                                    uintptr
	dllWkeNetSetData                                       uintptr
	dllWkeNetSetHTTPHeaderField                            uintptr
	dllWkeNetSetMIMEType                                   uintptr
	dllWkeNetStartUrlRequest                               uintptr
	dllWkeNodeOnCreateProcess                              uintptr
	dllWkeOnAlertBox                                       uintptr
	dllWkeOnConfirmBox                                     uintptr
	dllWkeOnConsole                                        uintptr
	dllWkeOnContextMenuItemClick                           uintptr
	dllWkeOnCreateView                                     uintptr
	dllWkeOnDidCreateScriptContext                         uintptr
	dllWkeOnDocumentReady                                  uintptr
	dllWkeOnDocumentReady2                                 uintptr
	dllWkeOnDownload                                       uintptr
	dllWkeOnDownload2                                      uintptr
	dllWkeOnDraggableRegionsChanged                        uintptr
	dllWkeOnLoadUrlBegin                                   uintptr
	dllWkeOnLoadUrlEnd                                     uintptr
	dllWkeOnLoadUrlFail                                    uintptr
	dllWkeOnLoadingFinish                                  uintptr
	dllWkeOnMouseOverUrlChanged                            uintptr
	dllWkeOnNavigation                                     uintptr
	dllWkeOnOtherLoad                                      uintptr
	dllWkeOnPaintBitUpdated                                uintptr
	dllWkeOnPaintUpdated                                   uintptr
	dllWkeOnPluginFind                                     uintptr
	dllWkeOnPrint                                          uintptr
	dllWkeOnPromptBox                                      uintptr
	dllWkeOnStartDragging                                  uintptr
	dllWkeOnTitleChanged                                   uintptr
	dllWkeOnURLChanged                                     uintptr
	dllWkeOnURLChanged2                                    uintptr
	dllWkeOnWillMediaLoad                                  uintptr
	dllWkeOnWillReleaseScriptContext                       uintptr
	dllWkeOnWindowClosing                                  uintptr
	dllWkeOnWindowDestroy                                  uintptr
	dllWkePaint                                            uintptr
	dllWkePaint2                                           uintptr
	dllWkePaste                                            uintptr
	dllWkePerformCookieCommand                             uintptr
	dllWkePluginListBuilderAddFileExtensionToLastMediaType uintptr
	dllWkePluginListBuilderAddMediaTypeToLastPlugin        uintptr
	dllWkePluginListBuilderAddPlugin                       uintptr
	dllWkePostURL                                          uintptr
	dllWkePostURLW                                         uintptr
	dllWkePrintToBitmap                                    uintptr
	dllWkeRegisterEmbedderCustomElement                    uintptr
	dllWkeReload                                           uintptr
	dllWkeRepaintIfNeeded                                  uintptr
	dllWkeResize                                           uintptr
	dllWkeResizeWindow                                     uintptr
	dllWkeRunJS                                            uintptr
	dllWkeRunJSW                                           uintptr
	dllWkeRunJsByFrame                                     uintptr
	dllWkeRunMessageLoop                                   uintptr
	dllWkeSaveMemoryCache                                  uintptr
	dllWkeScreenshot                                       uintptr
	dllWkeSelectAll                                        uintptr
	dllWkeSetClientHandler                                 uintptr
	dllWkeSetContextMenuEnabled                            uintptr
	dllWkeSetContextMenuItemShow                           uintptr
	dllWkeSetCookie                                        uintptr
	dllWkeSetCookieEnabled                                 uintptr
	dllWkeSetCookieJarFullPath                             uintptr
	dllWkeSetCookieJarPath                                 uintptr
	dllWkeSetCspCheckEnable                                uintptr
	dllWkeSetCursorInfoType                                uintptr
	dllWkeSetDebugConfig                                   uintptr
	dllWkeSetDeviceParameter                               uintptr
	dllWkeSetDirty                                         uintptr
	dllWkeSetDragDropEnable                                uintptr
	dllWkeSetDragEnable                                    uintptr
	dllWkeSetDragFiles                                     uintptr
	dllWkeSetEditable                                      uintptr
	dllWkeSetFileSystem                                    uintptr
	dllWkeSetFocus                                         uintptr
	dllWkeSetHandle                                        uintptr
	dllWkeSetHandleOffset                                  uintptr
	dllWkeSetHeadlessEnabled                               uintptr
	dllWkeSetLanguage                                      uintptr
	dllWkeSetLocalStorageFullPath                          uintptr
	dllWkeSetMediaPlayerFactory                            uintptr
	dllWkeSetMediaVolume                                   uintptr
	dllWkeSetMemoryCacheEnable                             uintptr
	dllWkeSetMouseEnabled                                  uintptr
	dllWkeSetName                                          uintptr
	dllWkeSetNavigationToNewWindowEnable                   uintptr
	dllWkeSetNpapiPluginsEnabled                           uintptr
	dllWkeSetProxy                                         uintptr
	dllWkeSetResourceGc                                    uintptr
	dllWkeSetString                                        uintptr
	dllWkeSetStringW                                       uintptr
	dllWkeSetStringWithoutNullTermination                  uintptr
	dllWkeSetTouchEnabled                                  uintptr
	dllWkeSetTransparent                                   uintptr
	dllWkeSetUIThreadCallback                              uintptr
	dllWkeSetUserAgent                                     uintptr
	dllWkeSetUserAgentW                                    uintptr
	dllWkeSetUserKeyValue                                  uintptr
	dllWkeSetViewNetInterface                              uintptr
	dllWkeSetViewProxy                                     uintptr
	dllWkeSetViewSettings                                  uintptr
	dllWkeSetWebViewName                                   uintptr
	dllWkeSetWindowTitle                                   uintptr
	dllWkeSetWindowTitleW                                  uintptr
	dllWkeSetZoomFactor                                    uintptr
	dllWkeShowDevtools                                     uintptr
	dllWkeShowWindow                                       uintptr
	dllWkeShutdown                                         uintptr
	dllWkeShutdownForDebug                                 uintptr
	dllWkeSleep                                            uintptr
	dllWkeStopLoading                                      uintptr
	dllWkeTitle                                            uintptr
	dllWkeTitleW                                           uintptr
	dllWkeToString                                         uintptr
	dllWkeToStringW                                        uintptr
	dllWkeUnfocus                                          uintptr
	dllWkeUpdate                                           uintptr
	dllWkeUtilBase64Decode                                 uintptr
	dllWkeUtilBase64Encode                                 uintptr
	dllWkeUtilCreateV8Snapshot                             uintptr
	dllWkeUtilDecodeURLEscape                              uintptr
	dllWkeUtilEncodeURLEscape                              uintptr
	dllWkeUtilPrintToPdf                                   uintptr
	dllWkeUtilRelasePrintPdfDatas                          uintptr
	dllWkeUtilSerializeToMHTML                             uintptr
	dllWkeUtilSetUiCallback                                uintptr
	dllWkeVersion                                          uintptr
	dllWkeVersionString                                    uintptr
	dllWkeVisitAllCookie                                   uintptr
	dllWkeWake                                             uintptr
	dllWkeWebFrameGetMainFrame                             uintptr
	dllWkeWebFrameGetMainWorldScriptContext                uintptr
	dllWkeWebViewName                                      uintptr
	dllWkeWidth                                            uintptr
	dllWkeZoomFactor                                       uintptr
)
