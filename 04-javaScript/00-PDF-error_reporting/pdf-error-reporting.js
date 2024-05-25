"use strict";
// Copyright (C) Microsoft Corporation. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
function hookErrorReporting(component, reportErrorCallback) {
	window.onerror = (message, source, lineno, columnNumber, error) => {
		const errorInfo = {
			component,
			name: error.name,
			line: lineno,
			column: columnNumber,
			source_url: source,
			stack: error.stack,
		};
		reportErrorCallback(errorInfo);
	};
	const consoleError = console.error;
	console.error = function (message, ...subst) {
		let stringifiedMessage;
		if (typeof message === 'string') {
			stringifiedMessage = message;
			//Log parameters for messages;
			// console.error('Expecte %d but got "%s", 5, "five")
			if (subst.length) {
				stringifiedMessage += ' | Params: ' + JSON.stringify(subst);
			}
		}
		else {
			try {
				stringifiedMessage = JSON.stringify(message);
			}
			catch (e) {
				stringifiedMessage = `JSON.stringify failed - ${String(message)}`;
			}
		}
		const errorInfo = {
			component,
			name: 'Console.Error',
			message: stringifiedMessage,
		};
		reportErrorCallback(errorInfo);
		consoleError(message, ...subst);
	};
}
