#!/usr/bin/env swift

import ApplicationServices
import Cocoa

func checkAccessibilityPermissions() -> Bool {
    let options = [kAXTrustedCheckOptionPrompt.takeUnretainedValue() as String: true]
    return AXIsProcessTrustedWithOptions(options as CFDictionary)
}

func getWindowByTitleFilter(id: String) -> AXUIElement? {
    guard checkAccessibilityPermissions() else { return nil }

    for app in NSWorkspace.shared.runningApplications {
        guard app.activationPolicy == .regular else { continue }

        let appElement = AXUIElementCreateApplication(app.processIdentifier)
        var windowsRef: AnyObject?
        guard
            AXUIElementCopyAttributeValue(appElement, kAXWindowsAttribute as CFString, &windowsRef)
                == .success,
            let windows = windowsRef as? [AXUIElement]
        else { continue }

        for window in windows {
            var titleRef: AnyObject?
            if AXUIElementCopyAttributeValue(window, kAXTitleAttribute as CFString, &titleRef) != .success {
                continue
            }
            guard let title = titleRef as? String else { continue }

            // `!title.contains("center")` is a hack to avoid running on the calling window.
            if title.contains(id) && !title.contains("center") {
                return window
            }
        }
    }
    return nil
}

func centerWindow(_ window: AXUIElement) -> Bool {
    guard let screen = NSScreen.main else { return false }

    // Set window size to 800x600
    var windowSize = CGSize(width: 900, height: 500)
    guard let sizeValue = AXValueCreate(.cgSize, &windowSize) else { return false }
    _ = AXUIElementSetAttributeValue(window, kAXSizeAttribute as CFString, sizeValue)

    let screenFrame = screen.visibleFrame
    var centerPosition = CGPoint(
        x: screenFrame.origin.x + (screenFrame.width - windowSize.width) / 2,
        y: screenFrame.origin.y + (screenFrame.height - windowSize.height) / 2
    )

    guard let positionValue = AXValueCreate(.cgPoint, &centerPosition) else { return false }
    return AXUIElementSetAttributeValue(window, kAXPositionAttribute as CFString, positionValue)
        == .success
}

guard CommandLine.arguments.count > 1 else {
    print("Usage: center.swift <window-title-filter>")
    exit(1)
}

if let window = getWindowByTitleFilter(id: CommandLine.arguments[1]) {
    _ = centerWindow(window)
} else {
    print("No window found")
    exit(1)
}
