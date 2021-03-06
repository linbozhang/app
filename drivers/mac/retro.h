#ifndef retro_h
#define retro_h

#if MAC_OS_X_VERSION_MAX_ALLOWED < 101200 

#import <Cocoa/Cocoa.h>
/* image.h */
static const NSCompositingOperation NSCompositingOperationCopy = NSCompositeCopy;

/* window.h */
#define NSWindowStyleMaskResizable NSResizableWindowMask
#define NSEventTypeLeftMouseDown NSLeftMouseDown
#define NSEventTypeLeftMouseUp NSLeftMouseUp
#define NSEventTypeRightMouseDown NSRightMouseDown
#define NSEventTypeRightMouseUp NSRightMouseUp
#define NSEventTypeOtherMouseDown NSOtherMouseDown
#define NSEventTypeOtherMouseUp NSOtherMouseUp
#define NSEventTypeScrollWheel NSScrollWheel
#define NSEventTypeMouseMoved NSMouseMoved
#define NSEventTypeLeftMouseDragged NSLeftMouseDragged
#define NSEventTypeRightMouseDragged NSRightMouseDragged
#define NSEventTypeOtherMouseDragged NSOtherMouseDragged
#define NSCompositingOperationCopy NSCompositeCopy
#define NSCompositingOperationSourceIn NSCompositeSourceIn
#define NSEventTypeFlagsChanged NSFlagsChanged
#define NSWindowStyleMaskTitled NSTitledWindowMask
#define NSWindowStyleMaskClosable NSClosableWindowMask
#define NSWindowStyleMaskMiniaturizable NSMiniaturizableWindowMask
#define NSWindowStyleMaskBorderless NSBorderlessWindowMask
#define NSWindowStyleMaskFullSizeContentView NSBorderlessWindowMask

/* menu.h */
#define NSEventModifierFlagControl NSControlKeyMask
#define NSEventModifierFlagShift NSShiftKeyMask
#define NSEventModifierFlagOption NSAlternateKeyMask
#define NSEventModifierFlagCommand NSCommandKeyMask
#define NSEventModifierFlagFunction NSFunctionKeyMask

#endif

#endif /* retro_h */