// Code generated by cmd/codegen from https://github.com/AllenDang/cimgui-go.
// DO NOT EDIT.

package imgui

// #include "extra_types.h"
// #include "cimnodes_structs_accessor.h"
// #include "cimnodes_wrapper.h"
import "C"

func NewEmulateThreeButtonMouse() *EmulateThreeButtonMouse {
	return newEmulateThreeButtonMouseFromC(C.EmulateThreeButtonMouse_EmulateThreeButtonMouse())
}

func (self *EmulateThreeButtonMouse) Destroy() {
	selfArg, selfFin := self.handle()
	C.EmulateThreeButtonMouse_destroy(selfArg)

	selfFin()
}

func NewNodesIO() *NodesIO {
	return newNodesIOFromC(C.ImNodesIO_ImNodesIO())
}

func (self *NodesIO) Destroy() {
	selfArg, selfFin := self.handle()
	C.ImNodesIO_destroy(selfArg)

	selfFin()
}

func NewNodesStyle() *NodesStyle {
	return newNodesStyleFromC(C.ImNodesStyle_ImNodesStyle())
}

func (self *NodesStyle) Destroy() {
	selfArg, selfFin := self.handle()
	C.ImNodesStyle_destroy(selfArg)

	selfFin()
}

func NewLinkDetachWithModifierClick() *LinkDetachWithModifierClick {
	return newLinkDetachWithModifierClickFromC(C.LinkDetachWithModifierClick_LinkDetachWithModifierClick())
}

func (self *LinkDetachWithModifierClick) Destroy() {
	selfArg, selfFin := self.handle()
	C.LinkDetachWithModifierClick_destroy(selfArg)

	selfFin()
}

func NewMultipleSelectModifier() *MultipleSelectModifier {
	return newMultipleSelectModifierFromC(C.MultipleSelectModifier_MultipleSelectModifier())
}

func (self *MultipleSelectModifier) Destroy() {
	selfArg, selfFin := self.handle()
	C.MultipleSelectModifier_destroy(selfArg)

	selfFin()
}

// ImNodesBeginInputAttributeV parameter default value hint:
// shape: ImNodesPinShape_CircleFilled
func ImNodesBeginInputAttributeV(id int32, shape NodesPinShape) {
	C.imnodes_BeginInputAttribute(C.int(id), C.ImNodesPinShape(shape))
}

func ImNodesBeginNode(id int32) {
	C.imnodes_BeginNode(C.int(id))
}

func ImNodesBeginNodeEditor() {
	C.imnodes_BeginNodeEditor()
}

func ImNodesBeginNodeTitleBar() {
	C.imnodes_BeginNodeTitleBar()
}

// ImNodesBeginOutputAttributeV parameter default value hint:
// shape: ImNodesPinShape_CircleFilled
func ImNodesBeginOutputAttributeV(id int32, shape NodesPinShape) {
	C.imnodes_BeginOutputAttribute(C.int(id), C.ImNodesPinShape(shape))
}

func ImNodesBeginStaticAttribute(id int32) {
	C.imnodes_BeginStaticAttribute(C.int(id))
}

func ImNodesClearLinkSelectionInt(link_id int32) {
	C.imnodes_ClearLinkSelection_Int(C.int(link_id))
}

func ImNodesClearLinkSelection() {
	C.imnodes_ClearLinkSelection_Nil()
}

func ImNodesClearNodeSelectionInt(node_id int32) {
	C.imnodes_ClearNodeSelection_Int(C.int(node_id))
}

func ImNodesClearNodeSelection() {
	C.imnodes_ClearNodeSelection_Nil()
}

func ImNodesCreateContext() *NodesContext {
	return newNodesContextFromC(C.imnodes_CreateContext())
}

// ImNodesDestroyContextV parameter default value hint:
// ctx: NULL
func ImNodesDestroyContextV(ctx *NodesContext) {
	ctxArg, ctxFin := ctx.handle()
	C.imnodes_DestroyContext(ctxArg)

	ctxFin()
}

func ImNodesEditorContextCreate() *NodesEditorContext {
	return newNodesEditorContextFromC(C.imnodes_EditorContextCreate())
}

func ImNodesEditorContextFree(noname1 *NodesEditorContext) {
	noname1Arg, noname1Fin := noname1.handle()
	C.imnodes_EditorContextFree(noname1Arg)

	noname1Fin()
}

func ImNodesEditorContextGetPanning() Vec2 {
	pOut := new(Vec2)
	pOutArg, pOutFin := wrap[C.ImVec2, *Vec2](pOut)

	C.imnodes_EditorContextGetPanning(pOutArg)

	pOutFin()

	return *pOut
}

func ImNodesEditorContextMoveToNode(node_id int32) {
	C.imnodes_EditorContextMoveToNode(C.int(node_id))
}

func ImNodesEditorContextResetPanning(pos Vec2) {
	C.imnodes_EditorContextResetPanning(pos.toC())
}

func ImNodesEditorContextSet(noname1 *NodesEditorContext) {
	noname1Arg, noname1Fin := noname1.handle()
	C.imnodes_EditorContextSet(noname1Arg)

	noname1Fin()
}

func ImNodesEndInputAttribute() {
	C.imnodes_EndInputAttribute()
}

func ImNodesEndNode() {
	C.imnodes_EndNode()
}

func ImNodesEndNodeEditor() {
	C.imnodes_EndNodeEditor()
}

func ImNodesEndNodeTitleBar() {
	C.imnodes_EndNodeTitleBar()
}

func ImNodesEndOutputAttribute() {
	C.imnodes_EndOutputAttribute()
}

func ImNodesEndStaticAttribute() {
	C.imnodes_EndStaticAttribute()
}

func ImNodesGetCurrentContext() *NodesContext {
	return newNodesContextFromC(C.imnodes_GetCurrentContext())
}

func ImNodesGetIO() *NodesIO {
	return newNodesIOFromC(C.imnodes_GetIO())
}

func ImNodesGetNodeDimensions(id int32) Vec2 {
	pOut := new(Vec2)
	pOutArg, pOutFin := wrap[C.ImVec2, *Vec2](pOut)

	C.imnodes_GetNodeDimensions(pOutArg, C.int(id))

	pOutFin()

	return *pOut
}

func ImNodesGetNodeEditorSpacePos(node_id int32) Vec2 {
	pOut := new(Vec2)
	pOutArg, pOutFin := wrap[C.ImVec2, *Vec2](pOut)

	C.imnodes_GetNodeEditorSpacePos(pOutArg, C.int(node_id))

	pOutFin()

	return *pOut
}

func ImNodesGetNodeGridSpacePos(node_id int32) Vec2 {
	pOut := new(Vec2)
	pOutArg, pOutFin := wrap[C.ImVec2, *Vec2](pOut)

	C.imnodes_GetNodeGridSpacePos(pOutArg, C.int(node_id))

	pOutFin()

	return *pOut
}

func ImNodesGetNodeScreenSpacePos(node_id int32) Vec2 {
	pOut := new(Vec2)
	pOutArg, pOutFin := wrap[C.ImVec2, *Vec2](pOut)

	C.imnodes_GetNodeScreenSpacePos(pOutArg, C.int(node_id))

	pOutFin()

	return *pOut
}

func ImNodesGetSelectedLinks(link_ids *int32) {
	link_idsArg, link_idsFin := WrapNumberPtr[C.int, int32](link_ids)
	C.imnodes_GetSelectedLinks(link_idsArg)

	link_idsFin()
}

func ImNodesGetSelectedNodes(node_ids *int32) {
	node_idsArg, node_idsFin := WrapNumberPtr[C.int, int32](node_ids)
	C.imnodes_GetSelectedNodes(node_idsArg)

	node_idsFin()
}

func ImNodesGetStyle() *NodesStyle {
	return newNodesStyleFromC(C.imnodes_GetStyle())
}

// ImNodesIsAnyAttributeActiveV parameter default value hint:
// attribute_id: NULL
func ImNodesIsAnyAttributeActiveV(attribute_id *int32) bool {
	attribute_idArg, attribute_idFin := WrapNumberPtr[C.int, int32](attribute_id)

	defer func() {
		attribute_idFin()
	}()
	return C.imnodes_IsAnyAttributeActive(attribute_idArg) == C.bool(true)
}

func ImNodesIsAttributeActive() bool {
	return C.imnodes_IsAttributeActive() == C.bool(true)
}

func ImNodesIsEditorHovered() bool {
	return C.imnodes_IsEditorHovered() == C.bool(true)
}

// ImNodesIsLinkCreatedBoolPtrV parameter default value hint:
// created_from_snap: NULL
func ImNodesIsLinkCreatedBoolPtrV(started_at_attribute_id *int32, ended_at_attribute_id *int32, created_from_snap *bool) bool {
	started_at_attribute_idArg, started_at_attribute_idFin := WrapNumberPtr[C.int, int32](started_at_attribute_id)
	ended_at_attribute_idArg, ended_at_attribute_idFin := WrapNumberPtr[C.int, int32](ended_at_attribute_id)
	created_from_snapArg, created_from_snapFin := WrapBool(created_from_snap)

	defer func() {
		started_at_attribute_idFin()
		ended_at_attribute_idFin()
		created_from_snapFin()
	}()
	return C.imnodes_IsLinkCreated_BoolPtr(started_at_attribute_idArg, ended_at_attribute_idArg, created_from_snapArg) == C.bool(true)
}

// ImNodesIsLinkCreatedIntPtrV parameter default value hint:
// created_from_snap: NULL
func ImNodesIsLinkCreatedIntPtrV(started_at_node_id *int32, started_at_attribute_id *int32, ended_at_node_id *int32, ended_at_attribute_id *int32, created_from_snap *bool) bool {
	started_at_node_idArg, started_at_node_idFin := WrapNumberPtr[C.int, int32](started_at_node_id)
	started_at_attribute_idArg, started_at_attribute_idFin := WrapNumberPtr[C.int, int32](started_at_attribute_id)
	ended_at_node_idArg, ended_at_node_idFin := WrapNumberPtr[C.int, int32](ended_at_node_id)
	ended_at_attribute_idArg, ended_at_attribute_idFin := WrapNumberPtr[C.int, int32](ended_at_attribute_id)
	created_from_snapArg, created_from_snapFin := WrapBool(created_from_snap)

	defer func() {
		started_at_node_idFin()
		started_at_attribute_idFin()
		ended_at_node_idFin()
		ended_at_attribute_idFin()
		created_from_snapFin()
	}()
	return C.imnodes_IsLinkCreated_IntPtr(started_at_node_idArg, started_at_attribute_idArg, ended_at_node_idArg, ended_at_attribute_idArg, created_from_snapArg) == C.bool(true)
}

func ImNodesIsLinkDestroyed(link_id *int32) bool {
	link_idArg, link_idFin := WrapNumberPtr[C.int, int32](link_id)

	defer func() {
		link_idFin()
	}()
	return C.imnodes_IsLinkDestroyed(link_idArg) == C.bool(true)
}

// ImNodesIsLinkDroppedV parameter default value hint:
// started_at_attribute_id: NULL
// including_detached_links: true
func ImNodesIsLinkDroppedV(started_at_attribute_id *int32, including_detached_links bool) bool {
	started_at_attribute_idArg, started_at_attribute_idFin := WrapNumberPtr[C.int, int32](started_at_attribute_id)

	defer func() {
		started_at_attribute_idFin()
	}()
	return C.imnodes_IsLinkDropped(started_at_attribute_idArg, C.bool(including_detached_links)) == C.bool(true)
}

func ImNodesIsLinkHovered(link_id *int32) bool {
	link_idArg, link_idFin := WrapNumberPtr[C.int, int32](link_id)

	defer func() {
		link_idFin()
	}()
	return C.imnodes_IsLinkHovered(link_idArg) == C.bool(true)
}

func ImNodesIsLinkSelected(link_id int32) bool {
	return C.imnodes_IsLinkSelected(C.int(link_id)) == C.bool(true)
}

func ImNodesIsLinkStarted(started_at_attribute_id *int32) bool {
	started_at_attribute_idArg, started_at_attribute_idFin := WrapNumberPtr[C.int, int32](started_at_attribute_id)

	defer func() {
		started_at_attribute_idFin()
	}()
	return C.imnodes_IsLinkStarted(started_at_attribute_idArg) == C.bool(true)
}

func ImNodesIsNodeHovered(node_id *int32) bool {
	node_idArg, node_idFin := WrapNumberPtr[C.int, int32](node_id)

	defer func() {
		node_idFin()
	}()
	return C.imnodes_IsNodeHovered(node_idArg) == C.bool(true)
}

func ImNodesIsNodeSelected(node_id int32) bool {
	return C.imnodes_IsNodeSelected(C.int(node_id)) == C.bool(true)
}

func ImNodesIsPinHovered(attribute_id *int32) bool {
	attribute_idArg, attribute_idFin := WrapNumberPtr[C.int, int32](attribute_id)

	defer func() {
		attribute_idFin()
	}()
	return C.imnodes_IsPinHovered(attribute_idArg) == C.bool(true)
}

func ImNodesLink(id int32, start_attribute_id int32, end_attribute_id int32) {
	C.imnodes_Link(C.int(id), C.int(start_attribute_id), C.int(end_attribute_id))
}

func ImNodesLoadCurrentEditorStateFromIniFile(file_name string) {
	file_nameArg, file_nameFin := WrapString(file_name)
	C.imnodes_LoadCurrentEditorStateFromIniFile(file_nameArg)

	file_nameFin()
}

func ImNodesLoadCurrentEditorStateFromIniString(data string, data_size uint64) {
	dataArg, dataFin := WrapString(data)
	C.imnodes_LoadCurrentEditorStateFromIniString(dataArg, C.xulong(data_size))

	dataFin()
}

func ImNodesLoadEditorStateFromIniFile(editor *NodesEditorContext, file_name string) {
	editorArg, editorFin := editor.handle()
	file_nameArg, file_nameFin := WrapString(file_name)
	C.imnodes_LoadEditorStateFromIniFile(editorArg, file_nameArg)

	editorFin()
	file_nameFin()
}

func ImNodesLoadEditorStateFromIniString(editor *NodesEditorContext, data string, data_size uint64) {
	editorArg, editorFin := editor.handle()
	dataArg, dataFin := WrapString(data)
	C.imnodes_LoadEditorStateFromIniString(editorArg, dataArg, C.xulong(data_size))

	editorFin()
	dataFin()
}

func ImNodesNumSelectedLinks() int32 {
	return int32(C.imnodes_NumSelectedLinks())
}

func ImNodesNumSelectedNodes() int32 {
	return int32(C.imnodes_NumSelectedNodes())
}

func ImNodesPopAttributeFlag() {
	C.imnodes_PopAttributeFlag()
}

func ImNodesPopColorStyle() {
	C.imnodes_PopColorStyle()
}

// ImNodesPopStyleVarV parameter default value hint:
// count: 1
func ImNodesPopStyleVarV(count int32) {
	C.imnodes_PopStyleVar(C.int(count))
}

func ImNodesPushAttributeFlag(flag NodesAttributeFlags) {
	C.imnodes_PushAttributeFlag(C.ImNodesAttributeFlags(flag))
}

func ImNodesPushColorStyle(item NodesCol, color uint32) {
	C.imnodes_PushColorStyle(C.ImNodesCol(item), C.uint(color))
}

func ImNodesPushStyleVarFloat(style_item NodesStyleVar, value float32) {
	C.imnodes_PushStyleVar_Float(C.ImNodesStyleVar(style_item), C.float(value))
}

func ImNodesPushStyleVarVec2(style_item NodesStyleVar, value Vec2) {
	C.imnodes_PushStyleVar_Vec2(C.ImNodesStyleVar(style_item), value.toC())
}

func ImNodesSaveCurrentEditorStateToIniFile(file_name string) {
	file_nameArg, file_nameFin := WrapString(file_name)
	C.imnodes_SaveCurrentEditorStateToIniFile(file_nameArg)

	file_nameFin()
}

// ImNodesSaveCurrentEditorStateToIniStringV parameter default value hint:
// data_size: NULL
func ImNodesSaveCurrentEditorStateToIniStringV(data_size *uint64) string {
	return C.GoString(C.imnodes_SaveCurrentEditorStateToIniString((*C.xulong)(data_size)))
}

func ImNodesSaveEditorStateToIniFile(editor *NodesEditorContext, file_name string) {
	editorArg, editorFin := editor.handle()
	file_nameArg, file_nameFin := WrapString(file_name)
	C.imnodes_SaveEditorStateToIniFile(editorArg, file_nameArg)

	editorFin()
	file_nameFin()
}

// ImNodesSaveEditorStateToIniStringV parameter default value hint:
// data_size: NULL
func ImNodesSaveEditorStateToIniStringV(editor *NodesEditorContext, data_size *uint64) string {
	editorArg, editorFin := editor.handle()

	defer func() {
		editorFin()
	}()
	return C.GoString(C.imnodes_SaveEditorStateToIniString(editorArg, (*C.xulong)(data_size)))
}

func ImNodesSelectLink(link_id int32) {
	C.imnodes_SelectLink(C.int(link_id))
}

func ImNodesSelectNode(node_id int32) {
	C.imnodes_SelectNode(C.int(node_id))
}

func ImNodesSetCurrentContext(ctx *NodesContext) {
	ctxArg, ctxFin := ctx.handle()
	C.imnodes_SetCurrentContext(ctxArg)

	ctxFin()
}

func ImNodesSetImGuiContext(ctx *Context) {
	ctxArg, ctxFin := ctx.handle()
	C.imnodes_SetImGuiContext(ctxArg)

	ctxFin()
}

func ImNodesSetNodeDraggable(node_id int32, draggable bool) {
	C.imnodes_SetNodeDraggable(C.int(node_id), C.bool(draggable))
}

func ImNodesSetNodeEditorSpacePos(node_id int32, editor_space_pos Vec2) {
	C.imnodes_SetNodeEditorSpacePos(C.int(node_id), editor_space_pos.toC())
}

func ImNodesSetNodeGridSpacePos(node_id int32, grid_pos Vec2) {
	C.imnodes_SetNodeGridSpacePos(C.int(node_id), grid_pos.toC())
}

func ImNodesSetNodeScreenSpacePos(node_id int32, screen_space_pos Vec2) {
	C.imnodes_SetNodeScreenSpacePos(C.int(node_id), screen_space_pos.toC())
}

func ImNodesSnapNodeToGrid(node_id int32) {
	C.imnodes_SnapNodeToGrid(C.int(node_id))
}

// ImNodesStyleColorsClassicV parameter default value hint:
// dest: NULL
func ImNodesStyleColorsClassicV(dest *NodesStyle) {
	destArg, destFin := dest.handle()
	C.imnodes_StyleColorsClassic(destArg)

	destFin()
}

// ImNodesStyleColorsDarkV parameter default value hint:
// dest: NULL
func ImNodesStyleColorsDarkV(dest *NodesStyle) {
	destArg, destFin := dest.handle()
	C.imnodes_StyleColorsDark(destArg)

	destFin()
}

// ImNodesStyleColorsLightV parameter default value hint:
// dest: NULL
func ImNodesStyleColorsLightV(dest *NodesStyle) {
	destArg, destFin := dest.handle()
	C.imnodes_StyleColorsLight(destArg)

	destFin()
}

func ImNodesBeginInputAttribute(id int32) {
	C.wrap_imnodes_BeginInputAttribute(C.int(id))
}

func ImNodesBeginOutputAttribute(id int32) {
	C.wrap_imnodes_BeginOutputAttribute(C.int(id))
}

func ImNodesDestroyContext() {
	C.wrap_imnodes_DestroyContext()
}

func ImNodesIsAnyAttributeActive() bool {
	return C.wrap_imnodes_IsAnyAttributeActive() == C.bool(true)
}

func ImNodesIsLinkCreatedBoolPtr(started_at_attribute_id *int32, ended_at_attribute_id *int32) bool {
	started_at_attribute_idArg, started_at_attribute_idFin := WrapNumberPtr[C.int, int32](started_at_attribute_id)
	ended_at_attribute_idArg, ended_at_attribute_idFin := WrapNumberPtr[C.int, int32](ended_at_attribute_id)

	defer func() {
		started_at_attribute_idFin()
		ended_at_attribute_idFin()
	}()
	return C.wrap_imnodes_IsLinkCreated_BoolPtr(started_at_attribute_idArg, ended_at_attribute_idArg) == C.bool(true)
}

func ImNodesIsLinkCreatedIntPtr(started_at_node_id *int32, started_at_attribute_id *int32, ended_at_node_id *int32, ended_at_attribute_id *int32) bool {
	started_at_node_idArg, started_at_node_idFin := WrapNumberPtr[C.int, int32](started_at_node_id)
	started_at_attribute_idArg, started_at_attribute_idFin := WrapNumberPtr[C.int, int32](started_at_attribute_id)
	ended_at_node_idArg, ended_at_node_idFin := WrapNumberPtr[C.int, int32](ended_at_node_id)
	ended_at_attribute_idArg, ended_at_attribute_idFin := WrapNumberPtr[C.int, int32](ended_at_attribute_id)

	defer func() {
		started_at_node_idFin()
		started_at_attribute_idFin()
		ended_at_node_idFin()
		ended_at_attribute_idFin()
	}()
	return C.wrap_imnodes_IsLinkCreated_IntPtr(started_at_node_idArg, started_at_attribute_idArg, ended_at_node_idArg, ended_at_attribute_idArg) == C.bool(true)
}

func ImNodesIsLinkDropped() bool {
	return C.wrap_imnodes_IsLinkDropped() == C.bool(true)
}

func ImNodesMiniMap() {
	C.wrap_imnodes_MiniMap()
}

func ImNodesPopStyleVar() {
	C.wrap_imnodes_PopStyleVar()
}

func ImNodesSaveCurrentEditorStateToIniString() string {
	return C.GoString(C.wrap_imnodes_SaveCurrentEditorStateToIniString())
}

func ImNodesSaveEditorStateToIniString(editor *NodesEditorContext) string {
	editorArg, editorFin := editor.handle()

	defer func() {
		editorFin()
	}()
	return C.GoString(C.wrap_imnodes_SaveEditorStateToIniString(editorArg))
}

func ImNodesStyleColorsClassic() {
	C.wrap_imnodes_StyleColorsClassic()
}

func ImNodesStyleColorsDark() {
	C.wrap_imnodes_StyleColorsDark()
}

func ImNodesStyleColorsLight() {
	C.wrap_imnodes_StyleColorsLight()
}

func (self NodesIO) SetEmulateThreeButtonMouse(v EmulateThreeButtonMouse) {
	vArg, _ := v.c()

	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesIO_SetEmulateThreeButtonMouse(selfArg, vArg)
}

func (self *NodesIO) EmulateThreeButtonMouse() EmulateThreeButtonMouse {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return *newEmulateThreeButtonMouseFromC(func() *C.EmulateThreeButtonMouse {
		result := C.wrap_ImNodesIO_GetEmulateThreeButtonMouse(selfArg)
		return &result
	}())
}

func (self NodesIO) SetLinkDetachWithModifierClick(v LinkDetachWithModifierClick) {
	vArg, _ := v.c()

	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesIO_SetLinkDetachWithModifierClick(selfArg, vArg)
}

func (self *NodesIO) LinkDetachWithModifierClick() LinkDetachWithModifierClick {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return *newLinkDetachWithModifierClickFromC(func() *C.LinkDetachWithModifierClick {
		result := C.wrap_ImNodesIO_GetLinkDetachWithModifierClick(selfArg)
		return &result
	}())
}

func (self NodesIO) SetMultipleSelectModifier(v MultipleSelectModifier) {
	vArg, _ := v.c()

	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesIO_SetMultipleSelectModifier(selfArg, vArg)
}

func (self *NodesIO) MultipleSelectModifier() MultipleSelectModifier {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return *newMultipleSelectModifierFromC(func() *C.MultipleSelectModifier {
		result := C.wrap_ImNodesIO_GetMultipleSelectModifier(selfArg)
		return &result
	}())
}

func (self NodesIO) SetAltMouseButton(v int32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesIO_SetAltMouseButton(selfArg, C.int(v))
}

func (self *NodesIO) AltMouseButton() int32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_ImNodesIO_GetAltMouseButton(selfArg))
}

func (self NodesIO) SetAutoPanningSpeed(v float32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesIO_SetAutoPanningSpeed(selfArg, C.float(v))
}

func (self *NodesIO) AutoPanningSpeed() float32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesIO_GetAutoPanningSpeed(selfArg))
}

func (self NodesStyle) SetGridSpacing(v float32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetGridSpacing(selfArg, C.float(v))
}

func (self *NodesStyle) GridSpacing() float32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetGridSpacing(selfArg))
}

func (self NodesStyle) SetNodeCornerRounding(v float32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetNodeCornerRounding(selfArg, C.float(v))
}

func (self *NodesStyle) NodeCornerRounding() float32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetNodeCornerRounding(selfArg))
}

func (self NodesStyle) SetNodePadding(v Vec2) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetNodePadding(selfArg, v.toC())
}

func (self *NodesStyle) NodePadding() Vec2 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return *(&Vec2{}).fromC(C.wrap_ImNodesStyle_GetNodePadding(selfArg))
}

func (self NodesStyle) SetNodeBorderThickness(v float32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetNodeBorderThickness(selfArg, C.float(v))
}

func (self *NodesStyle) NodeBorderThickness() float32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetNodeBorderThickness(selfArg))
}

func (self NodesStyle) SetLinkThickness(v float32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetLinkThickness(selfArg, C.float(v))
}

func (self *NodesStyle) LinkThickness() float32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetLinkThickness(selfArg))
}

func (self NodesStyle) SetLinkLineSegmentsPerLength(v float32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetLinkLineSegmentsPerLength(selfArg, C.float(v))
}

func (self *NodesStyle) LinkLineSegmentsPerLength() float32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetLinkLineSegmentsPerLength(selfArg))
}

func (self NodesStyle) SetLinkHoverDistance(v float32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetLinkHoverDistance(selfArg, C.float(v))
}

func (self *NodesStyle) LinkHoverDistance() float32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetLinkHoverDistance(selfArg))
}

func (self NodesStyle) SetPinCircleRadius(v float32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetPinCircleRadius(selfArg, C.float(v))
}

func (self *NodesStyle) PinCircleRadius() float32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetPinCircleRadius(selfArg))
}

func (self NodesStyle) SetPinQuadSideLength(v float32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetPinQuadSideLength(selfArg, C.float(v))
}

func (self *NodesStyle) PinQuadSideLength() float32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetPinQuadSideLength(selfArg))
}

func (self NodesStyle) SetPinTriangleSideLength(v float32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetPinTriangleSideLength(selfArg, C.float(v))
}

func (self *NodesStyle) PinTriangleSideLength() float32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetPinTriangleSideLength(selfArg))
}

func (self NodesStyle) SetPinLineThickness(v float32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetPinLineThickness(selfArg, C.float(v))
}

func (self *NodesStyle) PinLineThickness() float32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetPinLineThickness(selfArg))
}

func (self NodesStyle) SetPinHoverRadius(v float32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetPinHoverRadius(selfArg, C.float(v))
}

func (self *NodesStyle) PinHoverRadius() float32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetPinHoverRadius(selfArg))
}

func (self NodesStyle) SetPinOffset(v float32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetPinOffset(selfArg, C.float(v))
}

func (self *NodesStyle) PinOffset() float32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetPinOffset(selfArg))
}

func (self NodesStyle) SetMiniMapPadding(v Vec2) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetMiniMapPadding(selfArg, v.toC())
}

func (self *NodesStyle) MiniMapPadding() Vec2 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return *(&Vec2{}).fromC(C.wrap_ImNodesStyle_GetMiniMapPadding(selfArg))
}

func (self NodesStyle) SetMiniMapOffset(v Vec2) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetMiniMapOffset(selfArg, v.toC())
}

func (self *NodesStyle) MiniMapOffset() Vec2 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return *(&Vec2{}).fromC(C.wrap_ImNodesStyle_GetMiniMapOffset(selfArg))
}

func (self NodesStyle) SetFlags(v NodesStyleFlags) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetFlags(selfArg, C.ImNodesStyleFlags(v))
}

func (self *NodesStyle) Flags() NodesStyleFlags {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return NodesStyleFlags(C.wrap_ImNodesStyle_GetFlags(selfArg))
}

func (self NodesStyle) SetColors(v *[29]uint32) {
	vArg := make([]C.uint, len(v))
	for i, vV := range v {
		vArg[i] = C.uint(vV)
	}

	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetColors(selfArg, (*C.uint)(&vArg[0]))

	for i, vV := range vArg {
		(*v)[i] = uint32(vV)
	}
}

func (self *NodesStyle) Colors() [29]uint32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return func() [29]uint32 {
		result := [29]uint32{}
		resultMirr := C.wrap_ImNodesStyle_GetColors(selfArg)
		for i := range result {
			result[i] = uint32(C.cimnodes_unsigned_int_GetAtIdx(resultMirr, C.int(i)))
		}

		return result
	}()
}