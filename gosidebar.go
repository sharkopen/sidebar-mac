package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework CoreServices
#import <Foundation/Foundation.h>
#import <CoreServices/CoreServices.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

void addToSidebar(const char *fPath) {
	NSString *folderPath = [NSString stringWithUTF8String:fPath];
    NSURL *url = [NSURL fileURLWithPath:folderPath];
    LSSharedFileListRef listRef = LSSharedFileListCreate(NULL, kLSSharedFileListFavoriteItems, NULL);
    if (listRef == NULL) {
        NSLog(@"Error creating shared file list");
        return;
    }

    LSSharedFileListItemRef itemRef = LSSharedFileListInsertItemURL(listRef, kLSSharedFileListItemLast, NULL, NULL, (__bridge CFURLRef)url, NULL, NULL);
    if (itemRef == NULL) {
        NSLog(@"Error adding folder to sidebar");
        CFRelease(listRef);
        return;
    }

    CFRelease(listRef);
    NSLog(@"Folder added to sidebar successfully!");
}

void removeFromSidebar(const char *fPath) {
	NSString *folderPath = [NSString stringWithUTF8String:fPath];
    NSURL *url = [NSURL fileURLWithPath:folderPath];
    LSSharedFileListRef listRef = LSSharedFileListCreate(NULL, kLSSharedFileListFavoriteItems, NULL);
    if (listRef == NULL) {
        NSLog(@"Error creating shared file list");
        return;
    }

    NSArray *listSnapshot = (__bridge_transfer NSArray *)LSSharedFileListCopySnapshot(listRef, NULL);
    for (id item in listSnapshot) {
        LSSharedFileListItemRef itemRef = (__bridge LSSharedFileListItemRef)item;
        CFURLRef itemURLRef = NULL;
        if (LSSharedFileListItemResolve(itemRef, kLSSharedFileListNoUserInteraction | kLSSharedFileListDoNotMountVolumes, &itemURLRef, NULL) == noErr) {
            NSURL *itemURL = (__bridge_transfer NSURL *)itemURLRef;
			const char *urlCString = [itemURL.absoluteString UTF8String];
			printf("%s\n", urlCString);
			const char *folderUrl = [url.absoluteString UTF8String];
			printf("%s\n", folderUrl);
			size_t len = strlen(urlCString);
			char str = '/';
			char* result = malloc(len + 2); // Allocate memory for the result string
			strcpy(result, urlCString); // Copy the original string
			result[len] = str; // Append the character
			result[len + 1] = '\0'; // Null-terminate the string
            if (strcmp(folderUrl, result) == 0) {
               LSSharedFileListItemRemove(listRef, itemRef);
               NSLog(@"Folder removed from sidebar successfully!");
               break;
            }
			//if ([itemURL isEqual:url]) {
            //    LSSharedFileListItemRemove(listRef, itemRef);
            //    NSLog(@"Folder removed from sidebar successfully!");
            //    break;
            //}
        }
    }

    CFRelease(listRef);
}

void listSidebarItems() {
	LSSharedFileListRef listRef = LSSharedFileListCreate(NULL, kLSSharedFileListFavoriteItems, NULL);
    if (listRef == NULL) {
        NSLog(@"Error creating shared file list");
        return;
    }
    NSArray *listSnapshot = (__bridge_transfer NSArray *)LSSharedFileListCopySnapshot(listRef, NULL);
    for (id item in listSnapshot) {
        LSSharedFileListItemRef itemRef = (__bridge LSSharedFileListItemRef)item;
        CFURLRef itemURLRef = NULL;
        if (LSSharedFileListItemResolve(itemRef, kLSSharedFileListNoUserInteraction | kLSSharedFileListDoNotMountVolumes, &itemURLRef, NULL) == noErr) {
			NSURL *itemURL = (__bridge_transfer NSURL *)itemURLRef;
			const char *urlCString = [itemURL.absoluteString UTF8String];
			printf("%s\n", urlCString);

			//CFStringRef pathString = CFURLCopyFileSystemPath(itemURLRef, kCFURLPOSIXPathStyle);
			//const char* cString = CFStringGetCStringPtr(pathString, kCFStringEncodingUTF8);
			//printf("%s\n", cString);
			//CFRelease(pathString);
        }
    }
    CFRelease(listRef);
}

*/
import "C"
import (
	"fmt"
	"os"
	"unsafe"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("gosidebar [[add|rm] /path/to/folder, list]\n")
		return
	}
	if len(args) < 2 || len(args) > 3 {
		fmt.Printf("command is invalid\n")
		fmt.Printf("correct syntax is: gosidebar [[add|rm] /path/to/folder, list]\n")
		return
	}
	ops := args[1]

	if ops == "add" || ops == "rm" || ops == "list" {
		if ops == "list" {
			listSidebarItems()
			return
		}
		folderPath := args[2]
		_, err := os.Stat(folderPath)
		if os.IsNotExist(err) {
			fmt.Printf("folder not exist: %s, %s \n", folderPath, err.Error())
			return
		}
		if ops == "add" {
			// 添加文件夹到侧边栏
			addToSidebar(folderPath)
		} else if ops == "rm" {
			// 从侧边栏中移除文件夹
			removeFromSidebar(folderPath)
		} else {
			fmt.Printf("command invalid: %s \n", ops)
		}
	} else {
		fmt.Printf("the operation %s not support\n", ops)
	}
}

func addToSidebar(folderPath string) {
	cFolderPath := C.CString(folderPath)
	defer C.free(unsafe.Pointer(cFolderPath))

	C.addToSidebar(cFolderPath)
}

func removeFromSidebar(folderPath string) {
	cFolderPath := C.CString(folderPath)
	defer C.free(unsafe.Pointer(cFolderPath))

	C.removeFromSidebar(cFolderPath)
}

func listSidebarItems() {
	C.listSidebarItems()
}
