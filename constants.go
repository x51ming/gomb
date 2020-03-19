package gmb

type WkeDownloadOpt uintptr

const (
	KWkeDownloadOptCancel WkeDownloadOpt = iota
	KWkeDownloadOptCacheData
)

type WkeLoadingResult uint32

const (
	WKE_LOADING_SUCCEEDED WkeLoadingResult = iota
	WKE_LOADING_FAILED
	WKE_LOADING_CANCELED
)

type WkeRequestType uint32

const (
	KWkeRequestTypeInvalidation WkeRequestType = iota
	KWkeRequestTypeGet
	KWkeRequestTypePost
	KWkeRequestTypePut
)

type JsType uint32

const (
	JSTYPE_NUMBER JsType = iota
	JSTYPE_STRING
	JSTYPE_BOOLEAN
	JSTYPE_OBJECT
	JSTYPE_FUNCTION
	JSTYPE_UNDEFINED
	JSTYPE_ARRAY
	JSTYPE_NULL
)

type WkeOtherLoadType uint32

const (
	WKE_DID_START_LOADING WkeOtherLoadType = iota
	WKE_DID_STOP_LOADING
	WKE_DID_NAVIGATE
	WKE_DID_NAVIGATE_IN_PAGE
	WKE_DID_GET_RESPONSE_DETAILS
	WKE_DID_GET_REDIRECT_REQUEST
	WKE_DID_POST_REQUEST
)

type WkeResourceType uint32

const (
	WKE_RESOURCE_TYPE_MAIN_FRAME     WkeResourceType = iota // top level page
	WKE_RESOURCE_TYPE_SUB_FRAME                             // frame or iframe
	WKE_RESOURCE_TYPE_STYLESHEET                            // a CSS stylesheet
	WKE_RESOURCE_TYPE_SCRIPT                                // an external script
	WKE_RESOURCE_TYPE_IMAGE                                 // an image (jpg/gif/png/etc)
	WKE_RESOURCE_TYPE_FONT_RESOURCE                         // a font
	WKE_RESOURCE_TYPE_SUB_RESOURCE                          // an "other" subresource.
	WKE_RESOURCE_TYPE_OBJECT                                // an object (or embed) tag for a plugin,or a resource that a plugin requested.
	WKE_RESOURCE_TYPE_MEDIA                                 // a media resource.
	WKE_RESOURCE_TYPE_WORKER                                // the main resource of a dedicated worker.
	WKE_RESOURCE_TYPE_SHARED_WORKER                         // the main resource of a shared worker.
	WKE_RESOURCE_TYPE_PREFETCH                              // an explicitly requested prefetch
	WKE_RESOURCE_TYPE_FAVICON                               // a favicon
	WKE_RESOURCE_TYPE_XHR                                   // a XMLHttpRequest
	WKE_RESOURCE_TYPE_PING                                  // a ping request for <a ping>
	WKE_RESOURCE_TYPE_SERVICE_WORKER                        // the main resource of a service worker.
	WKE_RESOURCE_TYPE_LAST_TYPE
)

type WkeHttBodyElementType uint32

const (
	WkeHttBodyElementTypeData WkeHttBodyElementType = iota
	WkeHttBodyElementTypeFile
)

type WkeCookieCommand uint32

const (
	WkeCookieCommandClearAllCookies WkeCookieCommand = iota
	WkeCookieCommandClearSessionCookies
	WkeCookieCommandFlushCookiesToFile
	WkeCookieCommandReloadCookiesFromFile
)
