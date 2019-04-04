// Code generated by cdpgen. DO NOT EDIT.

package page

import (
	"github.com/mafredri/cdp/protocol/network"
	"github.com/mafredri/cdp/protocol/runtime"
	"github.com/mafredri/cdp/rpcc"
)

// DOMContentEventFiredClient is a client for DOMContentEventFired events.
type DOMContentEventFiredClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*DOMContentEventFiredReply, error)
	rpcc.Stream
}

// DOMContentEventFiredReply is the reply for DOMContentEventFired events.
type DOMContentEventFiredReply struct {
	Timestamp network.MonotonicTime `json:"timestamp"` // No description.
}

// FrameAttachedClient is a client for FrameAttached events. Fired when frame
// has been attached to its parent.
type FrameAttachedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*FrameAttachedReply, error)
	rpcc.Stream
}

// FrameAttachedReply is the reply for FrameAttached events.
type FrameAttachedReply struct {
	FrameID       FrameID             `json:"frameId"`         // Id of the frame that has been attached.
	ParentFrameID FrameID             `json:"parentFrameId"`   // Parent frame identifier.
	Stack         *runtime.StackTrace `json:"stack,omitempty"` // JavaScript stack trace of when frame was attached, only set if frame initiated from script.
}

// FrameClearedScheduledNavigationClient is a client for FrameClearedScheduledNavigation events.
// Fired when frame no longer has a scheduled navigation.
type FrameClearedScheduledNavigationClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*FrameClearedScheduledNavigationReply, error)
	rpcc.Stream
}

// FrameClearedScheduledNavigationReply is the reply for FrameClearedScheduledNavigation events.
type FrameClearedScheduledNavigationReply struct {
	FrameID FrameID `json:"frameId"` // Id of the frame that has cleared its scheduled navigation.
}

// FrameDetachedClient is a client for FrameDetached events. Fired when frame
// has been detached from its parent.
type FrameDetachedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*FrameDetachedReply, error)
	rpcc.Stream
}

// FrameDetachedReply is the reply for FrameDetached events.
type FrameDetachedReply struct {
	FrameID FrameID `json:"frameId"` // Id of the frame that has been detached.
}

// FrameNavigatedClient is a client for FrameNavigated events. Fired once
// navigation of the frame has completed. Frame is now associated with the new
// loader.
type FrameNavigatedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*FrameNavigatedReply, error)
	rpcc.Stream
}

// FrameNavigatedReply is the reply for FrameNavigated events.
type FrameNavigatedReply struct {
	Frame Frame `json:"frame"` // Frame object.
}

// FrameResizedClient is a client for FrameResized events.
type FrameResizedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*FrameResizedReply, error)
	rpcc.Stream
}

// FrameResizedReply is the reply for FrameResized events.
type FrameResizedReply struct {
}

// FrameRequestedNavigationClient is a client for FrameRequestedNavigation events.
// Fired when a renderer-initiated navigation is requested. Navigation may
// still be canceled after the event is issued.
type FrameRequestedNavigationClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*FrameRequestedNavigationReply, error)
	rpcc.Stream
}

// FrameRequestedNavigationReply is the reply for FrameRequestedNavigation events.
type FrameRequestedNavigationReply struct {
	FrameID FrameID                `json:"frameId"` // Id of the frame that has scheduled a navigation.
	Reason  ClientNavigationReason `json:"reason"`  // The reason for the navigation.
	URL     string                 `json:"url"`     // The destination URL for the requested navigation.
}

// FrameScheduledNavigationClient is a client for FrameScheduledNavigation events.
// Fired when frame schedules a potential navigation.
type FrameScheduledNavigationClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*FrameScheduledNavigationReply, error)
	rpcc.Stream
}

// FrameScheduledNavigationReply is the reply for FrameScheduledNavigation events.
type FrameScheduledNavigationReply struct {
	FrameID FrameID `json:"frameId"` // Id of the frame that has scheduled a navigation.
	Delay   float64 `json:"delay"`   // Delay (in seconds) until the navigation is scheduled to begin. The navigation is not guaranteed to start.
	// Reason The reason for the navigation.
	//
	// Values: "formSubmissionGet", "formSubmissionPost", "httpHeaderRefresh", "scriptInitiated", "metaTagRefresh", "pageBlockInterstitial", "reload".
	Reason string `json:"reason"`
	URL    string `json:"url"` // The destination URL for the scheduled navigation.
}

// FrameStartedLoadingClient is a client for FrameStartedLoading events. Fired
// when frame has started loading.
type FrameStartedLoadingClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*FrameStartedLoadingReply, error)
	rpcc.Stream
}

// FrameStartedLoadingReply is the reply for FrameStartedLoading events.
type FrameStartedLoadingReply struct {
	FrameID FrameID `json:"frameId"` // Id of the frame that has started loading.
}

// FrameStoppedLoadingClient is a client for FrameStoppedLoading events. Fired
// when frame has stopped loading.
type FrameStoppedLoadingClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*FrameStoppedLoadingReply, error)
	rpcc.Stream
}

// FrameStoppedLoadingReply is the reply for FrameStoppedLoading events.
type FrameStoppedLoadingReply struct {
	FrameID FrameID `json:"frameId"` // Id of the frame that has stopped loading.
}

// InterstitialHiddenClient is a client for InterstitialHidden events. Fired
// when interstitial page was hidden
type InterstitialHiddenClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*InterstitialHiddenReply, error)
	rpcc.Stream
}

// InterstitialHiddenReply is the reply for InterstitialHidden events.
type InterstitialHiddenReply struct {
}

// InterstitialShownClient is a client for InterstitialShown events. Fired
// when interstitial page was shown
type InterstitialShownClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*InterstitialShownReply, error)
	rpcc.Stream
}

// InterstitialShownReply is the reply for InterstitialShown events.
type InterstitialShownReply struct {
}

// JavascriptDialogClosedClient is a client for JavascriptDialogClosed events.
// Fired when a JavaScript initiated dialog (alert, confirm, prompt, or
// onbeforeunload) has been closed.
type JavascriptDialogClosedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*JavascriptDialogClosedReply, error)
	rpcc.Stream
}

// JavascriptDialogClosedReply is the reply for JavascriptDialogClosed events.
type JavascriptDialogClosedReply struct {
	Result    bool   `json:"result"`    // Whether dialog was confirmed.
	UserInput string `json:"userInput"` // User input in case of prompt.
}

// JavascriptDialogOpeningClient is a client for JavascriptDialogOpening events.
// Fired when a JavaScript initiated dialog (alert, confirm, prompt, or
// onbeforeunload) is about to open.
type JavascriptDialogOpeningClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*JavascriptDialogOpeningReply, error)
	rpcc.Stream
}

// JavascriptDialogOpeningReply is the reply for JavascriptDialogOpening events.
type JavascriptDialogOpeningReply struct {
	URL               string     `json:"url"`                     // Frame url.
	Message           string     `json:"message"`                 // Message that will be displayed by the dialog.
	Type              DialogType `json:"type"`                    // Dialog type.
	HasBrowserHandler bool       `json:"hasBrowserHandler"`       // True iff browser is capable showing or acting on the given dialog. When browser has no dialog handler for given target, calling alert while Page domain is engaged will stall the page execution. Execution can be resumed via calling Page.handleJavaScriptDialog.
	DefaultPrompt     *string    `json:"defaultPrompt,omitempty"` // Default dialog prompt.
}

// LifecycleEventClient is a client for LifecycleEvent events. Fired for top
// level page lifecycle events such as navigation, load, paint, etc.
type LifecycleEventClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*LifecycleEventReply, error)
	rpcc.Stream
}

// LifecycleEventReply is the reply for LifecycleEvent events.
type LifecycleEventReply struct {
	FrameID   FrameID               `json:"frameId"`   // Id of the frame.
	LoaderID  network.LoaderID      `json:"loaderId"`  // Loader identifier. Empty string if the request is fetched from worker.
	Name      string                `json:"name"`      // No description.
	Timestamp network.MonotonicTime `json:"timestamp"` // No description.
}

// LoadEventFiredClient is a client for LoadEventFired events.
type LoadEventFiredClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*LoadEventFiredReply, error)
	rpcc.Stream
}

// LoadEventFiredReply is the reply for LoadEventFired events.
type LoadEventFiredReply struct {
	Timestamp network.MonotonicTime `json:"timestamp"` // No description.
}

// NavigatedWithinDocumentClient is a client for NavigatedWithinDocument events.
// Fired when same-document navigation happens, e.g. due to history API usage
// or anchor navigation.
type NavigatedWithinDocumentClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*NavigatedWithinDocumentReply, error)
	rpcc.Stream
}

// NavigatedWithinDocumentReply is the reply for NavigatedWithinDocument events.
type NavigatedWithinDocumentReply struct {
	FrameID FrameID `json:"frameId"` // Id of the frame.
	URL     string  `json:"url"`     // Frame's new url.
}

// ScreencastFrameClient is a client for ScreencastFrame events. Compressed
// image data requested by the `startScreencast`.
type ScreencastFrameClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ScreencastFrameReply, error)
	rpcc.Stream
}

// ScreencastFrameReply is the reply for ScreencastFrame events.
type ScreencastFrameReply struct {
	Data      []byte                  `json:"data"`      // Base64-encoded compressed image.
	Metadata  ScreencastFrameMetadata `json:"metadata"`  // Screencast frame metadata.
	SessionID int                     `json:"sessionId"` // Frame number.
}

// ScreencastVisibilityChangedClient is a client for ScreencastVisibilityChanged events.
// Fired when the page with currently enabled screencast was shown or hidden `.
type ScreencastVisibilityChangedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ScreencastVisibilityChangedReply, error)
	rpcc.Stream
}

// ScreencastVisibilityChangedReply is the reply for ScreencastVisibilityChanged events.
type ScreencastVisibilityChangedReply struct {
	Visible bool `json:"visible"` // True if the page is visible.
}

// WindowOpenClient is a client for WindowOpen events. Fired when a new window
// is going to be opened, via window.open(), link click, form submission, etc.
type WindowOpenClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*WindowOpenReply, error)
	rpcc.Stream
}

// WindowOpenReply is the reply for WindowOpen events.
type WindowOpenReply struct {
	URL            string   `json:"url"`            // The URL for the new window.
	WindowName     string   `json:"windowName"`     // Window name.
	WindowFeatures []string `json:"windowFeatures"` // An array of enabled window features.
	UserGesture    bool     `json:"userGesture"`    // Whether or not it was triggered by user gesture.
}

// CompilationCacheProducedClient is a client for CompilationCacheProduced events.
// Issued for every compilation cache generated. Is only available if
// Page.setGenerateCompilationCache is enabled.
type CompilationCacheProducedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*CompilationCacheProducedReply, error)
	rpcc.Stream
}

// CompilationCacheProducedReply is the reply for CompilationCacheProduced events.
type CompilationCacheProducedReply struct {
	URL  string `json:"url"`  // No description.
	Data []byte `json:"data"` // Base64-encoded data
}
