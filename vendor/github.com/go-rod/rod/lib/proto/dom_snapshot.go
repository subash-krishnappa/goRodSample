// This file is generated by "./lib/proto/generate"

package proto

/*

DOMSnapshot

This domain facilitates obtaining document snapshots with DOM, layout, and style information.

*/

// DOMSnapshotDOMNode A Node in the DOM tree.
type DOMSnapshotDOMNode struct {
	// NodeType `Node`'s nodeType.
	NodeType int `json:"nodeType"`

	// NodeName `Node`'s nodeName.
	NodeName string `json:"nodeName"`

	// NodeValue `Node`'s nodeValue.
	NodeValue string `json:"nodeValue"`

	// TextValue (optional) Only set for textarea elements, contains the text value.
	TextValue string `json:"textValue,omitempty"`

	// InputValue (optional) Only set for input elements, contains the input's associated text value.
	InputValue string `json:"inputValue,omitempty"`

	// InputChecked (optional) Only set for radio and checkbox input elements, indicates if the element has been checked
	InputChecked bool `json:"inputChecked,omitempty"`

	// OptionSelected (optional) Only set for option elements, indicates if the element has been selected
	OptionSelected bool `json:"optionSelected,omitempty"`

	// BackendNodeID `Node`'s id, corresponds to DOM.Node.backendNodeId.
	BackendNodeID DOMBackendNodeID `json:"backendNodeId"`

	// ChildNodeIndexes (optional) The indexes of the node's child nodes in the `domNodes` array returned by `getSnapshot`, if
	// any.
	ChildNodeIndexes []int `json:"childNodeIndexes,omitempty"`

	// Attributes (optional) Attributes of an `Element` node.
	Attributes []*DOMSnapshotNameValue `json:"attributes,omitempty"`

	// PseudoElementIndexes (optional) Indexes of pseudo elements associated with this node in the `domNodes` array returned by
	// `getSnapshot`, if any.
	PseudoElementIndexes []int `json:"pseudoElementIndexes,omitempty"`

	// LayoutNodeIndex (optional) The index of the node's related layout tree node in the `layoutTreeNodes` array returned by
	// `getSnapshot`, if any.
	LayoutNodeIndex *int `json:"layoutNodeIndex,omitempty"`

	// DocumentURL (optional) Document URL that `Document` or `FrameOwner` node points to.
	DocumentURL string `json:"documentURL,omitempty"`

	// BaseURL (optional) Base URL that `Document` or `FrameOwner` node uses for URL completion.
	BaseURL string `json:"baseURL,omitempty"`

	// ContentLanguage (optional) Only set for documents, contains the document's content language.
	ContentLanguage string `json:"contentLanguage,omitempty"`

	// DocumentEncoding (optional) Only set for documents, contains the document's character set encoding.
	DocumentEncoding string `json:"documentEncoding,omitempty"`

	// PublicID (optional) `DocumentType` node's publicId.
	PublicID string `json:"publicId,omitempty"`

	// SystemID (optional) `DocumentType` node's systemId.
	SystemID string `json:"systemId,omitempty"`

	// FrameID (optional) Frame ID for frame owner elements and also for the document node.
	FrameID PageFrameID `json:"frameId,omitempty"`

	// ContentDocumentIndex (optional) The index of a frame owner element's content document in the `domNodes` array returned by
	// `getSnapshot`, if any.
	ContentDocumentIndex *int `json:"contentDocumentIndex,omitempty"`

	// PseudoType (optional) Type of a pseudo element node.
	PseudoType DOMPseudoType `json:"pseudoType,omitempty"`

	// ShadowRootType (optional) Shadow root type.
	ShadowRootType DOMShadowRootType `json:"shadowRootType,omitempty"`

	// IsClickable (optional) Whether this DOM node responds to mouse clicks. This includes nodes that have had click
	// event listeners attached via JavaScript as well as anchor tags that naturally navigate when
	// clicked.
	IsClickable bool `json:"isClickable,omitempty"`

	// EventListeners (optional) Details of the node's event listeners, if any.
	EventListeners []*DOMDebuggerEventListener `json:"eventListeners,omitempty"`

	// CurrentSourceURL (optional) The selected url for nodes with a srcset attribute.
	CurrentSourceURL string `json:"currentSourceURL,omitempty"`

	// OriginURL (optional) The url of the script (if any) that generates this node.
	OriginURL string `json:"originURL,omitempty"`

	// ScrollOffsetX (optional) Scroll offsets, set when this node is a Document.
	ScrollOffsetX *float64 `json:"scrollOffsetX,omitempty"`

	// ScrollOffsetY (optional) ...
	ScrollOffsetY *float64 `json:"scrollOffsetY,omitempty"`
}

// DOMSnapshotInlineTextBox Details of post layout rendered text positions. The exact layout should not be regarded as
// stable and may change between versions.
type DOMSnapshotInlineTextBox struct {
	// BoundingBox The bounding box in document coordinates. Note that scroll offset of the document is ignored.
	BoundingBox *DOMRect `json:"boundingBox"`

	// StartCharacterIndex The starting index in characters, for this post layout textbox substring. Characters that
	// would be represented as a surrogate pair in UTF-16 have length 2.
	StartCharacterIndex int `json:"startCharacterIndex"`

	// NumCharacters The number of characters in this post layout textbox substring. Characters that would be
	// represented as a surrogate pair in UTF-16 have length 2.
	NumCharacters int `json:"numCharacters"`
}

// DOMSnapshotLayoutTreeNode Details of an element in the DOM tree with a LayoutObject.
type DOMSnapshotLayoutTreeNode struct {
	// DomNodeIndex The index of the related DOM node in the `domNodes` array returned by `getSnapshot`.
	DomNodeIndex int `json:"domNodeIndex"`

	// BoundingBox The bounding box in document coordinates. Note that scroll offset of the document is ignored.
	BoundingBox *DOMRect `json:"boundingBox"`

	// LayoutText (optional) Contents of the LayoutText, if any.
	LayoutText string `json:"layoutText,omitempty"`

	// InlineTextNodes (optional) The post-layout inline text nodes, if any.
	InlineTextNodes []*DOMSnapshotInlineTextBox `json:"inlineTextNodes,omitempty"`

	// StyleIndex (optional) Index into the `computedStyles` array returned by `getSnapshot`.
	StyleIndex *int `json:"styleIndex,omitempty"`

	// PaintOrder (optional) Global paint order index, which is determined by the stacking order of the nodes. Nodes
	// that are painted together will have the same index. Only provided if includePaintOrder in
	// getSnapshot was true.
	PaintOrder *int `json:"paintOrder,omitempty"`

	// IsStackingContext (optional) Set to true to indicate the element begins a new stacking context.
	IsStackingContext bool `json:"isStackingContext,omitempty"`
}

// DOMSnapshotComputedStyle A subset of the full ComputedStyle as defined by the request whitelist.
type DOMSnapshotComputedStyle struct {
	// Properties Name/value pairs of computed style properties.
	Properties []*DOMSnapshotNameValue `json:"properties"`
}

// DOMSnapshotNameValue A name/value pair.
type DOMSnapshotNameValue struct {
	// Name Attribute/property name.
	Name string `json:"name"`

	// Value Attribute/property value.
	Value string `json:"value"`
}

// DOMSnapshotStringIndex Index of the string in the strings table.
type DOMSnapshotStringIndex int

// DOMSnapshotArrayOfStrings Index of the string in the strings table.
type DOMSnapshotArrayOfStrings []DOMSnapshotStringIndex

// DOMSnapshotRareStringData Data that is only present on rare nodes.
type DOMSnapshotRareStringData struct {
	// Index ...
	Index []int `json:"index"`

	// Value ...
	Value []DOMSnapshotStringIndex `json:"value"`
}

// DOMSnapshotRareBooleanData ...
type DOMSnapshotRareBooleanData struct {
	// Index ...
	Index []int `json:"index"`
}

// DOMSnapshotRareIntegerData ...
type DOMSnapshotRareIntegerData struct {
	// Index ...
	Index []int `json:"index"`

	// Value ...
	Value []int `json:"value"`
}

// DOMSnapshotRectangle ...
type DOMSnapshotRectangle []float64

// DOMSnapshotDocumentSnapshot Document snapshot.
type DOMSnapshotDocumentSnapshot struct {
	// DocumentURL Document URL that `Document` or `FrameOwner` node points to.
	DocumentURL DOMSnapshotStringIndex `json:"documentURL"`

	// Title Document title.
	Title DOMSnapshotStringIndex `json:"title"`

	// BaseURL Base URL that `Document` or `FrameOwner` node uses for URL completion.
	BaseURL DOMSnapshotStringIndex `json:"baseURL"`

	// ContentLanguage Contains the document's content language.
	ContentLanguage DOMSnapshotStringIndex `json:"contentLanguage"`

	// EncodingName Contains the document's character set encoding.
	EncodingName DOMSnapshotStringIndex `json:"encodingName"`

	// PublicID `DocumentType` node's publicId.
	PublicID DOMSnapshotStringIndex `json:"publicId"`

	// SystemID `DocumentType` node's systemId.
	SystemID DOMSnapshotStringIndex `json:"systemId"`

	// FrameID Frame ID for frame owner elements and also for the document node.
	FrameID DOMSnapshotStringIndex `json:"frameId"`

	// Nodes A table with dom nodes.
	Nodes *DOMSnapshotNodeTreeSnapshot `json:"nodes"`

	// Layout The nodes in the layout tree.
	Layout *DOMSnapshotLayoutTreeSnapshot `json:"layout"`

	// TextBoxes The post-layout inline text nodes.
	TextBoxes *DOMSnapshotTextBoxSnapshot `json:"textBoxes"`

	// ScrollOffsetX (optional) Horizontal scroll offset.
	ScrollOffsetX *float64 `json:"scrollOffsetX,omitempty"`

	// ScrollOffsetY (optional) Vertical scroll offset.
	ScrollOffsetY *float64 `json:"scrollOffsetY,omitempty"`

	// ContentWidth (optional) Document content width.
	ContentWidth *float64 `json:"contentWidth,omitempty"`

	// ContentHeight (optional) Document content height.
	ContentHeight *float64 `json:"contentHeight,omitempty"`
}

// DOMSnapshotNodeTreeSnapshot Table containing nodes.
type DOMSnapshotNodeTreeSnapshot struct {
	// ParentIndex (optional) Parent node index.
	ParentIndex []int `json:"parentIndex,omitempty"`

	// NodeType (optional) `Node`'s nodeType.
	NodeType []int `json:"nodeType,omitempty"`

	// ShadowRootType (optional) Type of the shadow root the `Node` is in. String values are equal to the `ShadowRootType` enum.
	ShadowRootType *DOMSnapshotRareStringData `json:"shadowRootType,omitempty"`

	// NodeName (optional) `Node`'s nodeName.
	NodeName []DOMSnapshotStringIndex `json:"nodeName,omitempty"`

	// NodeValue (optional) `Node`'s nodeValue.
	NodeValue []DOMSnapshotStringIndex `json:"nodeValue,omitempty"`

	// BackendNodeID (optional) `Node`'s id, corresponds to DOM.Node.backendNodeId.
	BackendNodeID []DOMBackendNodeID `json:"backendNodeId,omitempty"`

	// Attributes (optional) Attributes of an `Element` node. Flatten name, value pairs.
	Attributes []DOMSnapshotArrayOfStrings `json:"attributes,omitempty"`

	// TextValue (optional) Only set for textarea elements, contains the text value.
	TextValue *DOMSnapshotRareStringData `json:"textValue,omitempty"`

	// InputValue (optional) Only set for input elements, contains the input's associated text value.
	InputValue *DOMSnapshotRareStringData `json:"inputValue,omitempty"`

	// InputChecked (optional) Only set for radio and checkbox input elements, indicates if the element has been checked
	InputChecked *DOMSnapshotRareBooleanData `json:"inputChecked,omitempty"`

	// OptionSelected (optional) Only set for option elements, indicates if the element has been selected
	OptionSelected *DOMSnapshotRareBooleanData `json:"optionSelected,omitempty"`

	// ContentDocumentIndex (optional) The index of the document in the list of the snapshot documents.
	ContentDocumentIndex *DOMSnapshotRareIntegerData `json:"contentDocumentIndex,omitempty"`

	// PseudoType (optional) Type of a pseudo element node.
	PseudoType *DOMSnapshotRareStringData `json:"pseudoType,omitempty"`

	// PseudoIdentifier (optional) Pseudo element identifier for this node. Only present if there is a
	// valid pseudoType.
	PseudoIdentifier *DOMSnapshotRareStringData `json:"pseudoIdentifier,omitempty"`

	// IsClickable (optional) Whether this DOM node responds to mouse clicks. This includes nodes that have had click
	// event listeners attached via JavaScript as well as anchor tags that naturally navigate when
	// clicked.
	IsClickable *DOMSnapshotRareBooleanData `json:"isClickable,omitempty"`

	// CurrentSourceURL (optional) The selected url for nodes with a srcset attribute.
	CurrentSourceURL *DOMSnapshotRareStringData `json:"currentSourceURL,omitempty"`

	// OriginURL (optional) The url of the script (if any) that generates this node.
	OriginURL *DOMSnapshotRareStringData `json:"originURL,omitempty"`
}

// DOMSnapshotLayoutTreeSnapshot Table of details of an element in the DOM tree with a LayoutObject.
type DOMSnapshotLayoutTreeSnapshot struct {
	// NodeIndex Index of the corresponding node in the `NodeTreeSnapshot` array returned by `captureSnapshot`.
	NodeIndex []int `json:"nodeIndex"`

	// Styles Array of indexes specifying computed style strings, filtered according to the `computedStyles` parameter passed to `captureSnapshot`.
	Styles []DOMSnapshotArrayOfStrings `json:"styles"`

	// Bounds The absolute position bounding box.
	Bounds []DOMSnapshotRectangle `json:"bounds"`

	// Text Contents of the LayoutText, if any.
	Text []DOMSnapshotStringIndex `json:"text"`

	// StackingContexts Stacking context information.
	StackingContexts *DOMSnapshotRareBooleanData `json:"stackingContexts"`

	// PaintOrders (optional) Global paint order index, which is determined by the stacking order of the nodes. Nodes
	// that are painted together will have the same index. Only provided if includePaintOrder in
	// captureSnapshot was true.
	PaintOrders []int `json:"paintOrders,omitempty"`

	// OffsetRects (optional) The offset rect of nodes. Only available when includeDOMRects is set to true
	OffsetRects []DOMSnapshotRectangle `json:"offsetRects,omitempty"`

	// ScrollRects (optional) The scroll rect of nodes. Only available when includeDOMRects is set to true
	ScrollRects []DOMSnapshotRectangle `json:"scrollRects,omitempty"`

	// ClientRects (optional) The client rect of nodes. Only available when includeDOMRects is set to true
	ClientRects []DOMSnapshotRectangle `json:"clientRects,omitempty"`

	// BlendedBackgroundColors (experimental) (optional) The list of background colors that are blended with colors of overlapping elements.
	BlendedBackgroundColors []DOMSnapshotStringIndex `json:"blendedBackgroundColors,omitempty"`

	// TextColorOpacities (experimental) (optional) The list of computed text opacities.
	TextColorOpacities []float64 `json:"textColorOpacities,omitempty"`
}

// DOMSnapshotTextBoxSnapshot Table of details of the post layout rendered text positions. The exact layout should not be regarded as
// stable and may change between versions.
type DOMSnapshotTextBoxSnapshot struct {
	// LayoutIndex Index of the layout tree node that owns this box collection.
	LayoutIndex []int `json:"layoutIndex"`

	// Bounds The absolute position bounding box.
	Bounds []DOMSnapshotRectangle `json:"bounds"`

	// Start The starting index in characters, for this post layout textbox substring. Characters that
	// would be represented as a surrogate pair in UTF-16 have length 2.
	Start []int `json:"start"`

	// Length The number of characters in this post layout textbox substring. Characters that would be
	// represented as a surrogate pair in UTF-16 have length 2.
	Length []int `json:"length"`
}

// DOMSnapshotDisable Disables DOM snapshot agent for the given page.
type DOMSnapshotDisable struct{}

// ProtoReq name
func (m DOMSnapshotDisable) ProtoReq() string { return "DOMSnapshot.disable" }

// Call sends the request
func (m DOMSnapshotDisable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// DOMSnapshotEnable Enables DOM snapshot agent for the given page.
type DOMSnapshotEnable struct{}

// ProtoReq name
func (m DOMSnapshotEnable) ProtoReq() string { return "DOMSnapshot.enable" }

// Call sends the request
func (m DOMSnapshotEnable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// DOMSnapshotGetSnapshot (deprecated) Returns a document snapshot, including the full DOM tree of the root node (including iframes,
// template contents, and imported documents) in a flattened array, as well as layout and
// white-listed computed style information for the nodes. Shadow DOM in the returned DOM tree is
// flattened.
type DOMSnapshotGetSnapshot struct {
	// ComputedStyleWhitelist Whitelist of computed styles to return.
	ComputedStyleWhitelist []string `json:"computedStyleWhitelist"`

	// IncludeEventListeners (optional) Whether or not to retrieve details of DOM listeners (default false).
	IncludeEventListeners bool `json:"includeEventListeners,omitempty"`

	// IncludePaintOrder (optional) Whether to determine and include the paint order index of LayoutTreeNodes (default false).
	IncludePaintOrder bool `json:"includePaintOrder,omitempty"`

	// IncludeUserAgentShadowTree (optional) Whether to include UA shadow tree in the snapshot (default false).
	IncludeUserAgentShadowTree bool `json:"includeUserAgentShadowTree,omitempty"`
}

// ProtoReq name
func (m DOMSnapshotGetSnapshot) ProtoReq() string { return "DOMSnapshot.getSnapshot" }

// Call the request
func (m DOMSnapshotGetSnapshot) Call(c Client) (*DOMSnapshotGetSnapshotResult, error) {
	var res DOMSnapshotGetSnapshotResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// DOMSnapshotGetSnapshotResult (deprecated) ...
type DOMSnapshotGetSnapshotResult struct {
	// DomNodes The nodes in the DOM tree. The DOMNode at index 0 corresponds to the root document.
	DomNodes []*DOMSnapshotDOMNode `json:"domNodes"`

	// LayoutTreeNodes The nodes in the layout tree.
	LayoutTreeNodes []*DOMSnapshotLayoutTreeNode `json:"layoutTreeNodes"`

	// ComputedStyles Whitelisted ComputedStyle properties for each node in the layout tree.
	ComputedStyles []*DOMSnapshotComputedStyle `json:"computedStyles"`
}

// DOMSnapshotCaptureSnapshot Returns a document snapshot, including the full DOM tree of the root node (including iframes,
// template contents, and imported documents) in a flattened array, as well as layout and
// white-listed computed style information for the nodes. Shadow DOM in the returned DOM tree is
// flattened.
type DOMSnapshotCaptureSnapshot struct {
	// ComputedStyles Whitelist of computed styles to return.
	ComputedStyles []string `json:"computedStyles"`

	// IncludePaintOrder (optional) Whether to include layout object paint orders into the snapshot.
	IncludePaintOrder bool `json:"includePaintOrder,omitempty"`

	// IncludeDOMRects (optional) Whether to include DOM rectangles (offsetRects, clientRects, scrollRects) into the snapshot
	IncludeDOMRects bool `json:"includeDOMRects,omitempty"`

	// IncludeBlendedBackgroundColors (experimental) (optional) Whether to include blended background colors in the snapshot (default: false).
	// Blended background color is achieved by blending background colors of all elements
	// that overlap with the current element.
	IncludeBlendedBackgroundColors bool `json:"includeBlendedBackgroundColors,omitempty"`

	// IncludeTextColorOpacities (experimental) (optional) Whether to include text color opacity in the snapshot (default: false).
	// An element might have the opacity property set that affects the text color of the element.
	// The final text color opacity is computed based on the opacity of all overlapping elements.
	IncludeTextColorOpacities bool `json:"includeTextColorOpacities,omitempty"`
}

// ProtoReq name
func (m DOMSnapshotCaptureSnapshot) ProtoReq() string { return "DOMSnapshot.captureSnapshot" }

// Call the request
func (m DOMSnapshotCaptureSnapshot) Call(c Client) (*DOMSnapshotCaptureSnapshotResult, error) {
	var res DOMSnapshotCaptureSnapshotResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// DOMSnapshotCaptureSnapshotResult ...
type DOMSnapshotCaptureSnapshotResult struct {
	// Documents The nodes in the DOM tree. The DOMNode at index 0 corresponds to the root document.
	Documents []*DOMSnapshotDocumentSnapshot `json:"documents"`

	// Strings Shared string table that all string properties refer to with indexes.
	Strings []string `json:"strings"`
}
