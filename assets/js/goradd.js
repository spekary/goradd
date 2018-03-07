var $j, qcubed, qc;

(function( $ ) {

$j = $;

$j.fn.extend({
    wait: function(time, type) {
        time = time || 1000;
        type = type || "fx";
        return this.queue(type, function() {
            var self = this;

            setTimeout(function() {
                $j(self).dequeue();
            }, time);
        });
    }
});


/**
 * @namespace goradd
 */
goradd = {
    /**
     * Queued Ajax requests.
     * A new Ajax request won't be started until the previous queued
     * request has finished.
     * @param {function} o function that returns ajax options.
     * @param {boolean} blnAsync true to launch right away.
     */
    ajaxQueue: function(o, blnAsync) {
        if (typeof $j.ajaxq === "undefined" || blnAsync) {
            $j.ajax(o()); // fallback in case ajaxq is not here
        } else {
            $j.ajaxq("goradd", o);
        }
    },
    ajaxQueueIsRunning: function() {
        if ($j.ajaxq) {
            return $j.ajaxq.isRunning("goradd");
        }
        return false;
    },

    /**
     * Adds a value to the next ajax or server post for the specified control. You can either call this ongoing, or
     * call it in response to the "posting" event. This is the preferred way for custom javascript controls to send data
     * to their goradd counterparts.
     *
     * @param {string} strControlId
     * @param {string} strProperty
     * @param {mixed} strNewValue
     */
    setControlValue: function(strControlId, strProperty, strNewValue) {
        if (!gr.controlValues[strControlId]) {
            gr.controlValues[strControlId] = {};
        }
        gr.controlValues[strControlId][strProperty] = strNewValue;
    },
    /**
     * Given a control, returns the correct index to use in the formObjsModified array.
     * @param ctl
     * @private
     */
    _formObjChangeIndex: function (ctl) {
        var id = $j(ctl).attr('id'),
            strType = $j(ctl).prop("type"),
            name = $j(ctl).attr("name"),
            indexOffset;

        if (((strType === 'checkbox') || (strType === 'radio')) &&
           id && ((indexOffset = id.lastIndexOf('_')) >= 0)) { // a member of a control list
            return id.substr(0, indexOffset); // use the id of the group
        }
        else if (id && strType === 'radio' && name !== id) { // a radio button with a group name
            return id; // these buttons are changed individually
        }
        else if (id && strType === 'hidden') { // a hidden field, possibly associated with a different widget
            if ((indexOffset = id.lastIndexOf('_')) >= 0) {
                return id.substr(0, indexOffset); // use the id of the parent control
            }
            return name;
        }
        else if (name && !id) {
            name = name.replace('[]', ''); // remove brackets if they are there for array
            return name;
        }
        return id;
    },
    /**
     * Records that a control has changed in order to synchronize the control with
     * the php version on the next request.
     * @param event
     */
    formObjChanged: function (event) {
        var ctl = event.target,
            id = gr._formObjChangeIndex(ctl),
            strType = $j(ctl).prop("type"),
            name = $j(ctl).attr("name");

        if (strType === 'radio' && name !== id) { // a radio button with a group name
            // since html does not submit a changed event on the deselected radio, we are going to invalidate all the controls in the group
            var group = $j('input[name=' + name + ']');
            if (group) {
                group.each(function () {
                    id = $j(this).attr('id');
                    gr.formObjsModified[id] = true;
                });
            }
        }
        else if (id) {
            gr.formObjsModified[id] = true;
        }
    },
    /**
     * Initialize form related scripts
     * @param {string} strFormId
     */
    initForm: function () {
        var $form =  $j(goradd.getForm());
        $form.on ('formObjChanged', gr.formObjChanged); // Allow any control, including hidden inputs, to trigger a change and post of its data.
        $form.submit(function(event) {
            if (!$j('#Goradd__Params').val()) { // did postBack initiate the submit?
                // if not, prevent implicit form submission. This can happen in the rare case we have a single field and no submit button.
                event.preventDefault();
            }
        });
        goradd.registerControls();
    },

    /**
     * Post the form. ServerActions go here.
     *
     * @param {object} params An object containing the following:
     *  controlId {string}: The control id to post an action to
     *  eventId {int}: The event id
     *  async: If true, process the event asynchronously without waiting for other events to complete
     *  values {object} (optional): An optional object, that contains values coming to send with the event
     *      event: The event's action value, if one is provided. This can be any type, including an object.
     *      action: The action's action value, if one is provided. Any type.
     *      control: The control's action value, if one is provided. Any type.
     *
     */
    postBack: function(params) {
        if (gr.blockEvents) {
            return;  // We are waiting for a response from the server
        }

        var $objForm = $j(goradd.getForm());
        var formId = $objForm.attr("id");

        var checkableControls = $objForm.find('input[type="checkbox"], input[type="radio"]');
        params.checkableControls = gr._checkableControlValues(formId, $j.makeArray(checkableControls));

        params.callType = "Server";

        // Notify custom controls that we are about to post
        $objForm.trigger("posting", "Server");

        if (!$j.isEmptyObject(gr.controlValues)) {
            params.controlValues = gr.controlValues;
        }
        $j('#Goradd__Params').val(JSON.stringify(params));

        // have $ trigger the submit event (so it can catch all submit events)
        $objForm.trigger("submit");
    },
    /**
     * This function resolves the state of checkable controls into postable values.
     *
     * Checkable controls (checkboxes and radio buttons) can be problematic. We have the following issues to work around:
     * - On a submit, only the values of the checked items are submitted. Non-checked items are not submitted.
     * - QCubed may have checkboxes that are part of the form object, but not visible on the html page. In particular,
     *   this can happen when a grid is creating objects at render time, and then scrolls or pages so those objects
     *   are no longer "visible".
     * - Controls can be part of a group, and the group gets the value of the checked control(s), rather than individual
     *   items getting a true or false.
     *
     * To solve all of these issues, we post a value that has all the values of all visible checked items, either
     * true or false for individual items, or an array of values, single value, or null for groups. QCubed controls that
     * deal with checkable controls must look for this special posted variable to know how to update their internal state.
     *
     * Checkboxes that are part of a group will return an array of values, keyed by the group id.
     * Radio buttons that are part of a group will return a single value keyed by group id.
     * Checkboxes and radio buttons that are not part of a group will return a true or false keyed by the control id.
     * Note that for radio buttons, a group is defined by a common identifier in the id. Radio buttons with the same
     * name, but different ids, are not considered part of a group for purposes here, even though visually they will
     * act like they are part of a group. This allows you to create individual QRadioButton objects that each will
     * be updated with a true or false, but the browser will automatically make sure only one is checked.
     *
     * Any time an id has an underscore in it, that control is considered part of a group. The value after the underscore
     * will be the value returned, and before the last underscore will be id that will be used as the key for the value.
     *
     * @param {string} strForm   Form Id
     * @param {Array} controls  Array of checkable controls. These must be checkable controls, it will not validate this.
     * @returns {object}  A hash of values keyed by control id
     * @private
     */
    _checkableControlValues: function(strForm, controls) {
        var values = {};

        if (!controls || controls.length === 0) {
            return {};
        }
        $j.each(controls, function() {
            var $element = $j(this),
                id = $element.attr("id"),
                strType = $element.prop("type"),
                index = null,
                offset;

            if (id &&
                (offset = id.lastIndexOf('_')) !== -1) {
                // A control group
                index = id.substr(offset + 1);
                id = id.substr(0, offset);
            }
            switch (strType) {
                case "checkbox":
                    if (index !== null) {   // this is a group of checkboxes
                        var a = values[id];
                        if ($element.is(":checked")) {
                            if (a) {
                                a.push(index);
                            } else {
                                a = [index];
                            }
                            values[id] = a;
                        }
                        else {
                            if (!a) {
                                values[id] = null; // empty array to notify that the group has a null value, if nothing gets checked
                            }
                        }
                    } else {
                        values[id] = $element.is(":checked");
                    }
                    break;

                case "radio":
                    if (index !== null) {
                        if ($element.is(":checked")) {
                            values[id] = index;
                        }
                    } else {
                        // control name MIGHT be a group name, which we don't want here, so we use control id instead
                        values[id] = $element.is(":checked");
                    }
                    break;
            }
        });
        return values;
    },

    /**
     * Gets the data to be sent to an ajax call as post data. This will be called from the ajax queueing function, and
     * will erase the cache of changed objects to prepare for the next call.
     *
     * @param {object} params An object containing the following:
     *  controlId {string}: The control id to post an action to
     *  eventId {int}: The event id
     *  async: If true, process the event asynchronously without waiting for other events to complete
     *  formId: The id of the form getting posted
     *  values {object} (optional): An optional object, that contains values coming to send with the event
     *      event: The event's action value, if one is provided. This can be any type, including an object.
     *      action: The action's action value, if one is provided. Any type.
     *      control: The control's action value, if one is provided. Any type.
     * @return {object} Post Data
     */
    _getAjaxData: function(params) {
        var $form = $j('#' + params.formId),
            $formElements = $form.find('input,select,textarea'),
            checkables = [],
            controls = [],
            postData = {};

        // Notify controls we are about to post.
        $form.trigger("posting", "Ajax");

        // Filter and separate controls into checkable and non-checkable controls
        // We ignore controls that have not changed to reduce the amount of data sent in an ajax post.
        $formElements.each(function() {
            var $element = $j(this),
                id = $element.attr("id"),
                blnForm = (id && (id.substr(0, 8) === 'Goradd__')),
                strType = $element.prop("type"),
                objChangeIndex = gr._formObjChangeIndex($element);


                if (!gr.inputSupport || // if not oninput support, then post all the controls, rather than just the modified ones
                gr.ajaxError || // Ajax error would mean that formObjsModified is invalid. We need to submit everything.
                (objChangeIndex && gr.formObjsModified[objChangeIndex]) ||
                blnForm) {  // all controls with Goradd__ at the beginning of the id are always posted.

                switch (strType) {
                    case "checkbox":
                    case "radio":
                        checkables.push(this);
                        break;

                    default:
                        controls.push(this);
                }
            }
        });


        $j.each(controls, function() {
            var $element = $j(this),
                strType = $element.prop("type"),
                strControlId = $element.attr("id"),
                strControlName = $element.attr("name"),
                strPostValue = $element.val();
            var strPostName = (strControlName ? strControlName: strControlId);

            switch (strType) {
                case "select-multiple":
                    var items = $element.find(':selected'),
                        values = [];
                    if (items.length) {
                        values = $j.map($j.makeArray(items), function(item) {
                            return $j(item).val();
                        });
                        postData[strPostName] = values;
                    }
                    else {
                        postData[strPostName] = null;    // mark it as set to nothing
                    }
                    break;

                default:
                    postData[strPostName] = strPostValue;
                    break;
            }
        });

        // Update most of the Goradd__ parameters explicitly here. Others, like the state and form id will have been handled above.
        params.callType = "Ajax"
        if (!$j.isEmptyObject(gr.controlValues)) {
            params.controlValues = gr.controlValues;
        }
        params.checkableValues = gr._checkableControlValues(params.formId, checkables);
        postData.Goradd__Params = JSON.stringify(params);

        gr.ajaxError = false;
        gr.formObjsModified = {};
        gr.controlValues = {};

        return postData;
    },

    /**
     * @param {object} params An object containing the following:
     *  controlId {string}: The control id to post an action to
     *  eventId {int}: The event id
     *  async: If true, process the event asynchronously without waiting for other events to complete
     *  values {object} (optional): An optional object, that contains values coming to send with the event
     *      event: The event's action value, if one is provided. This can be any type, including an object.
     *      action: The action's action value, if one is provided. Any type.
     *      control: The control's action value, if one is provided. Any type.
     *
     * @return {void}
     */
    postAjax: function(params) {
        var $objForm = $j(goradd.getForm()),
            formAction = $objForm.attr("action"),
            async = params.hasOwnProperty("async");

        if (gr.blockEvents) {
            return;
        }

        params.formId = $objForm.attr('id');

        // Use an ajax queue so ajax requests happen synchronously
        gr.ajaxQueue(function() {
            var data = gr._getAjaxData(params);

            return {
                url: formAction,
                type: "POST",
                data: data,
                error: function (XMLHttpRequest, textStatus, errorThrown) {
                    var result = XMLHttpRequest.responseText;

                    if (XMLHttpRequest.status !== 0 || (result && result.length > 0)) {
                        gr.displayAjaxError(result, textStatus, errorThrown);
                        return false;
                    } else {
                        gr.displayAjaxError("Unknown ajax error", '', '');
                        return false;
                    }
                },
                success: function (json) {
                    gr._prevUpdateTime = new Date().getTime();
                    if ($j.type(json) === 'string') {
                        // If server has a problem sending any ajax response, like when headers are already sent, we will get that error as a string here
                        gr.displayAjaxError(json, '', '');
                        return false;
                    }
                    if (json.js) {
                        var deferreds = [];
                        // Load all javascript files before attempting to process the rest of the response, in case some things depend on the injected files
                        $j.each(json.js, function (i, v) {
                            deferreds.push(gr.loadJavaScriptFile(v));
                        });
                        gr.processImmediateAjaxResponse(json, params); // go ahead and begin processing things that will not depend on the javascript files to allow parallel processing
                        $j.when.apply($j, deferreds).then(
                            function () {
                                gr.processDeferredAjaxResponse(json);
                                gr.blockEvents = false;
                            }, // success
                            function () {
                                window.console.log('Failed to load a file');
                                gr.blockEvents = false;
                            } // failed to load a file. What to do?
                        );
                    } else {
                        gr.processImmediateAjaxResponse(json, params);
                        gr.processDeferredAjaxResponse(json);
                        gr.blockEvents = false;
                    }
                }
            };
        }, async);
    },
    displayAjaxError: function(resultText, textStatus, errorThrown) {
        var objErrorWindow;

        gr.ajaxError = true;
        gr.blockEvents = false;

        if (resultText.substr(0, 15) === '<!DOCTYPE html>') {
            window.alert("An error occurred.\r\n\r\nThe error response will appear in a new popup.");
            objErrorWindow = window.open('about:blank', 'qcubed_error', 'menubar=no,toolbar=no,location=no,status=no,scrollbars=yes,resizable=yes,width=1000,height=700,left=50,top=50');
            objErrorWindow.focus();
            objErrorWindow.document.write(resultText);
        } else {
            resultText = $j('<div>').html(resultText);
            $j('<div id="Goradd_AJAX_Error" />')
                .append('<h1 style="text-transform:capitalize">' + textStatus + '</h1>')
                .append('<p>' + errorThrown + '</p>')
                .append(resultText)
                .append('<button onclick="$j(this).parent().hide()">OK</button>')
                .appendTo('form');
        }
    },

    /**
     * Start me up.
     */
    initialize: function() {
        ////////////////////////////////
        // Browser-related functionality
        ////////////////////////////////

        gr.loadJavaScriptFile = function(strScript, objCallback) {
            return $j.ajax({
                url: strScript,
                success: objCallback,
                dataType: "script",
                cache: true
            });
        };

        gr.loadStyleSheetFile = function(strStyleSheetFile, strMediaType) {
            if (strMediaType){
                strMediaType = " media="+strMediaType;
            }
            $j('head').append('<link rel="stylesheet"'+strMediaType+' href="' + strStyleSheetFile + '" type="text/css" />');
        };

        /////////////////////////////
        // Form-related functionality
        /////////////////////////////
        
        $j(window).on ("storage", function (o) {
            if (o.originalEvent.key === "goradd.broadcast") {
                gr.updateForm();
            }
        });

        gr.inputSupport = 'oninput' in document;

        // Detect browsers that do not correctly support the oninput event, even though they say they do.
        // IE 9 in particular has a major bug
        var ua = window.navigator.userAgent;
        var intIeOffset = ua.indexOf ('MSIE');
        if (intIeOffset > -1) {
            var intOffset2 = ua.indexOf ('.', intIeOffset + 5);
            var strVersion = ua.substr (intIeOffset + 5, intOffset2 - intIeOffset - 5);
            if (strVersion < 10) {
                gr.inputSupport = false;
            }
        }

        $j( document ).ajaxComplete(function( event, request, settings ) {
            if (!gr.ajaxQueueIsRunning()) {
                gr.processFinalCommands();
            }
        });

        // TODO: Add a detector of the back button. This detector should ping the server to make sure the formstate exists on the server. If not,
        // it should reload the page.
        return this;
    },
    processImmediateAjaxResponse: function(json, params) {
        if (json.controls) {
            $j.each(json.controls, function() {
                var strControlId = this.id,
                    $control = $j(goradd.getControl(strControlId)),
                    $wrapper = $j(goradd.getWrapper(strControlId));

                if (this.value !== undefined) {
                    $control.val(this.value);
                }

                if (this.attributes !== undefined) {
                    $control.attr (this.attributes);
                }

                if (this.html !== undefined) {
                    if ($wrapper.length) {
                        // Control's wrapper was found, so fill it in
                        $wrapper.html(this.html);
                    }
                    else if ($control.length) {
                        // control was found without a wrapper, replace it in the same position it was in.
                        // remove related controls (error, name ...) for wrapper-less controls
                        var relSelector = "[data-qrel='" + strControlId + "']",
                            relItems = $j(relSelector),
                            $relParent;

                        if (relItems && relItems.length) {
                            // if the control is wrapped in a related control, we move the control outside the related controls
                            // before deleting the related controls
                            $relParent = $control.parents(relSelector).last();
                            if ($relParent.length) {
                                $control.insertBefore($relParent);
                            }
                            relItems.remove();
                        }

                        $control.before(this.html).remove();
                    }
                    else {
                        // control is being injected at the top level, so put it at the end of the form.
                        var $objForm = $j(goradd.getForm());
                        $objForm.append(this.html);
                    }
                }
            });
        }

        gr.registerControls();

        if (json.watcher && params.controlId) {
            gr.broadcastChange();
        }
        if (json.ss) {
            $j.each(json.ss, function (i,v) {
                gr.loadStyleSheetFile(v, "all");
            });
        }
        if (json.alert) {
            $j.each(json.alert, function (i,v) {
                window.alert(v);
            });
        }
    },
    processDeferredAjaxResponse: function(json) {
        if (json.commands) { // commands
            $j.each(json.commands, function (index, command) {
                if (command.final &&
                    gr.ajaxQueueIsRunning()) {

                    gr.enqueueFinalCommand(command);
                } else {
                    gr.processCommand(command);
                }
            });
        }
        if (json.winclose) {
            window.close();
        }
        if (json.loc) {
            if (json.loc === 'reload') {
                window.location.reload(true);
            } else {
                document.location = json.loc;
            }
        }
    },
    processCommand: function(command) {
        var params,
            objs;

        if (command.script) {
            eval (command.script);
        }
        else if (command.selector) {
            params = gr.unpackArray(command.params);

            if (typeof command.selector === 'string') {
                objs = $j(command.selector);
            } else {
                objs = $j(command.selector[0], command.selector[1]);
            }

            // apply the function on each jQuery object found, using the found jQuery object as the context.
            objs.each (function () {
                var $item = $j(this);
                if ($item[command.func]) {
                    $item[command.func].apply($j(this), params);
                }
            });
        }
        else if (command.func) {
            params = gr.unpackArray(command.params);

            // Find the function by name. Walk an object list in the process.
            objs = command.func.split(".");
            var obj = window;
            var ctx = null;

            $j.each (objs, function (i, v) {
                ctx = obj;
                obj = obj[v];
            });
            // obj is now a function object, and ctx is the parent of the function object
            obj.apply(ctx, params);
        }

    },
    enqueueFinalCommand: function(command) {
        gr.finalCommands.push(command);
    },
    processFinalCommands: function() {
        while(gr.finalCommands.length) {
            var command = gr.finalCommands.pop();
            gr.processCommand(command);
        }
    },
    /**
     * Convert from JSON return value to an actual jQuery object. Certain structures don't work in JSON, like closures,
     * but can be part of a javascript object.
     * @param params
     * @returns {*}
     */
    unpackArray: function(params) {
        if (!params) {
            return null;
        }
        var newParams = [];

        $j.each(params, function (index, item){
            if ($j.type(item) === 'object') {
                if (item.goraddObject) {
                    item = gr.unpackObj(item);  // top level special object
                }
                else {
                    // look for special objects inside top level objects.
                    var newItem = {};
                    $j.each (item, function (key, obj) {
                        newItem[key] = gr.unpackObj(obj);
                    });
                    item = newItem;
                }
            }
            else if ($j.type(item) === 'array') {
                item = gr.unpackArray (item);
            }
            newParams.push(item);
        });
        return newParams;
    },

    /**
     * Given an object coming from goradd, will attempt to decode the object into a corresponding javascript object.
     * @param obj
     * @returns {*}
     */
    unpackObj: function (obj) {
        if ($j.type(obj) === 'object' &&
                obj.goraddObject) {

            switch (obj.goraddObject) {
                case 'closure':
                    if (obj.params) {
                        params = [];
                        $j.each (obj.params, function (i, v) {
                            params.push(gr.unpackObj(v)); // recurse
                        });

                        return new Function(params, obj.func);
                    } else {
                        return new Function(obj.func);
                    }
                    break;

                case 'dateTime':
                    return new Date(obj.year, obj.month, obj.day, obj.hour, obj.minute, obj.second);

                case 'varName':
                    // Find the variable value starting at the window context.
                    var vars = obj.varName.split(".");
                    var val = window;
                    $j.each (vars, function (i, v) {
                        val = val[v];
                    });
                    return val;

                case 'func':
                    // Returns the result of the given function called immediately
                    // Find the function and context starting at the window context.
                    var target = window;
                    var params;
                    if (obj.context) {
                       var objects = obj.context.split(".");
                        $j.each (objects, function (i, v) {
                            target = target[v];
                        });
                    }

                    if (obj.params) {
                        params = [];
                        $j.each (obj.params, function (i, v) {
                            params.push(gr.unpackObj(v)); // recurse
                        });
                    }
                    var func = target[obj.func];

                    return func.apply(target, params);
            }
        }
        else if ($j.type(obj) === 'object') {
            var newItem = {};
            $j.each (obj, function (key, obj2) {
                newItem[key] = gr.unpackObj(obj2);
            });
            return newItem;
        }
        else if ($j.type(obj) === 'array') {
            return gr.unpackArray(obj);
        }
        return obj; // no change
    },
    setCookie: function(name, val, expires, path, dom, secure) {
            var cookie = name + "=" + encodeURIComponent(val) + "; ";

            if (expires) {
                cookie += "expires=" + expires.toUTCString() + "; ";
            }

            if (path) {
                cookie += "path=" + path + "; ";
            }
            if (dom) {
                cookie += "domain=" + dom + "; ";
            }
            if (secure) {
                cookie += "secure;";
            }

            document.cookie = cookie;
        }
};

///////////////////////////////
// Timers-related functionality
///////////////////////////////

goradd._objTimers = {};

goradd.clearTimeout = function(strTimerId) {
    if (goradd._objTimers[strTimerId]) {
        clearTimeout(goradd._objTimers[strTimerId]);
        goradd._objTimers[strTimerId] = null;
    }
};

goradd.setTimeout = function(strTimerId, action, intDelay) {
    goradd.clearTimeout(strTimerId);
    goradd._objTimers[strTimerId] = setTimeout(action, intDelay);
};

goradd.startTimer = function(strControlId, intDeltaTime, blnPeriodic) {
    var strTimerId = strControlId + '_ct';
    gr.stopTimer(strControlId, blnPeriodic);
    if (blnPeriodic) {
        goradd._objTimers[strTimerId] = setInterval(function() {
            $j('#' + strControlId).trigger('timerexpiredevent')
        }, intDeltaTime);
    } else {
        goradd._objTimers[strTimerId] = setTimeout(function() {
            $j('#' + strControlId).trigger('timerexpiredevent')
        }, intDeltaTime);
    }
};

goradd.stopTimer = function(strControlId, blnPeriodic) {
    var strTimerId = strControlId + '_ct';
    if (goradd._objTimers[strTimerId]) {
        if (blnPeriodic) {
            clearInterval(goradd._objTimers[strTimerId]);
        } else {
            clearTimeout(goradd._objTimers[strTimerId]);
        }
        goradd._objTimers[strTimerId] = null;
    }
};

///////////////////////////////
// Watcher support
///////////////////////////////
goradd._prevUpdateTime = 0;
goradd.minUpdateInterval = 1000; // milliseconds to limit broadcast updates. Feel free to change this.
goradd.broadcastChange = function () {
    if ('localStorage' in window && window.localStorage !== null) {
        var newTime = new Date().getTime();
        localStorage.setItem("goradd.broadcast", newTime); // must change value to induce storage event in other windows
    }
};

goradd.updateForm = function() {
    // call this whenever you generally just need the form to update without a specific action.
    var newTime = new Date().getTime();

    // the following code prevents too many updates from happening in a short amount of time.
    // the default will update no faster than once per second.
    if (newTime - goradd._prevUpdateTime >= goradd.minUpdateInterval) {
        //refresh immediately
        goradd.postAjax ({});
        goradd.clearTimeout ('goradd.update');
    } else if (!goradd._objTimers['goradd.update']) {
        // delay to let multiple fast actions only trigger periodic refreshes
        goradd.setTimeout ('goradd.update', 'goradd.updateForm', goradd.minUpdateInterval);
    }
};

/////////////////////////////////////
// Drag and drop support
/////////////////////////////////////

goradd.draggable = function (parentId, draggableId) {
    // we are working around some jQuery UI bugs here..
    $j('#' + parentId).on("dragstart", function () {
        var c = $j(this);
        c.data ("originalPosition", c.position());
    }).on("dragstop", function () {
        var c = $j(this);
        gr.setControlValue(draggableId, "_DragData", {originalPosition: {left: c.data("originalPosition").left, top: c.data("originalPosition").top}, position: {left: c.position().left, top: c.position().top}});
    });
};

goradd.droppable = function (parentId, droppableId) {
    $j('#' + parentId).on("drop", function (event, ui) {
        gr.setControlValue(droppableId, "_DroppedId", ui.draggable.attr("id"));
    });
};

goradd.resizable = function (parentId, resizeableId) {
    $j('#' + parentId).on("resizestart", function () {
        var c = $j(this);
        c.data ("oW", c.width());
        c.data ("oH", c.height());
    })
    .on("resizestop", function () {
        var c = $j(this);
        gr.setControlValue(resizeableId, "_ResizeData", {originalSize: {width: c.data("oW"), height: c.data("oH")} , size:{width: c.width(), height: c.height()}});
    });
};

/////////////////////////////////////
// JQueryUI Support
/////////////////////////////////////

goradd.dialog = function(controlId) {
    $j('#' + controlId).on ("keydown", "input,select", function(event) {
        // makes sure a return key fires the default button if there is one
        if (event.which === 13) {
            var b = $j(this).closest("[role=\'dialog\']").find("button[type=\'submit\']");
            if (b && b[0]) {
                b[0].click();
            }
            event.preventDefault();
        }
    });
};

goradd.accordion = function(controlId) {
    $j('#' + controlId).on("accordionactivate", function(event, ui) {
        goradd.setControlValue(controlId, "_SelectedIndex", $j(this).accordion("option", "active"));
        $j(this).trigger("change");
    });
};

goradd.progressbar = function(controlId) {
    $j('#' + controlId).on("progressbarchange", function (event, ui) {
        goradd.setControlValue(controlId, "_Value", $j(this).progressbar ("value"));
    });
};

goradd.selectable = function(controlId) {
    $j('#' + controlId).on("selectablestop", function (event, ui) {
        var strItems;

        strItems = "";
        $j(".ui-selected", this).each(function() {
            strItems = strItems + "," + this.id;
        });

        if (strItems) {
            strItems = strItems.substring (1);
        }
        goradd.setControlValue(controlId, "_SelectedItems", strItems);

    });
};

goradd.slider = function(controlId) {
    $j('#' + controlId).on("slidechange", function (event, ui) {
        if (ui.values && ui.values.length) {
            gr.setControlValue(controlId, "_Values", ui.values[0] + ',' +  ui.values[1]);
        } else {
            gr.setControlValue(controlId, "_Value", ui.value);
        }
    });
};

goradd.tabs = function(controlId) {
    $j('#' + controlId).on("tabsactivate", function(event, ui) {
        var i = $j(this).tabs( "option", "active" );
        var id = ui.newPanel ? ui.newPanel.attr("id") : null;
        gr.setControlValue(controlId, "_active", [i,id]);
    });
};

goradd.datagrid2 = function(controlId) {
    $j('#' + controlId).on("click", "thead tr th a", function(event, ui) {
        var cellIndex = $j(this).parent()[0].cellIndex;
        $j(this).trigger('qdg2sort', cellIndex); // Triggers the QDataGrid_SortEvent
        event.stopPropagation();
    });
};

goradd.dialog = function(controlId) {
    $j('#' + controlId).on("tabsactivate", function(event, ui) {
        var i = $j(this).tabs( "option", "active" );
        var id = ui.newPanel ? ui.newPanel.attr("id") : null;
        gr.setControlValue(controlId, "_active", [i,id]);
    });
};

/////////////////////////////////
// Controls-related functionality
/////////////////////////////////

goradd.getControl = function(controlId) {
    return document.getElementById(controlId);
};

goradd.getWrapper = function(mixControl) {
    if (typeof mixControl === 'string') {
        return document.getElementById(mixControl + "_ctl")
    }
    else {
        return document.getElementById($j(mixControl).attr('id') + "_ctl")
    }
};

goradd.getForm = function() {
    return $j('form[data-goradd="form"]')[0]
};

/**
 * Radio buttons are a little tricky to set if they are part of a group
 * @param strControlId
 */
goradd.setRadioInGroup = function(strControlId) {
    var $objControl = $j('#' + strControlId);
    if ($objControl) {
        var groupName = $objControl.prop('name');
        if (groupName) {
            var $radios = $objControl.closest('form').find('input[type=radio][name=' + groupName + ']');
            $radios.val([strControlId]);  // jquery does the work here of setting just the one control
            $radios.trigger('formObjChanged'); // send the new values back to the form
        }
    }
};

/////////////////////////////
// Register Control - General
/////////////////////////////

goradd.controlValues = {};
goradd.formObjsModified = {};
goradd.ajaxError = false;
goradd.inputSupport = true;
goradd.blockEvents = false;
goradd.finalCommands = [];

goradd.registerControl = function(mixControl) {
    var objControl = goradd.getControl(mixControl),
        objWrapper;

    if (!objControl) {
        return;
    }

    var $control = $j(objControl);

    if ($control.data('gr') === 'reg') {
        return // this control is already registered
    }

    // detect changes to objects before any changes trigger other events
    if (objControl.type === 'checkbox' || objControl.type === 'radio') {
        // clicks are equivalent to changes for checkboxes and radio buttons, but some browsers send change way after a click. We need to capture the click first.
        $j(objControl).on ('click', gr.formObjChanged);
    }
    $j(objControl).on ('change input', gr.formObjChanged);
    $j(objControl).on ('change input', 'input, select, textarea', gr.formObjChanged);   // make sure we get to bubbled events before later attached handlers


    // Link the Wrapper and the Control together
    objWrapper = gr.getWapper(objControl.id);
    if (objWrapper) {
        objWrapper.control = objControl;
    }
    $control.data('gr', 'reg') // mark the control as registered so we don't attach events twice
};

goradd.registerControls = function() {
    $j('[data-goradd="ctl"]').each(function() {
        goradd.registerControl(this);
    });
};

})( jQuery );

////////////////////////////////
// QCubed Shortcuts and Initialize
////////////////////////////////

gr = goradd;
gr.pB = gr.postBack;
gr.pA = gr.postAjax;
gr.getC = gr.getControl;
gr.getW = gr.getWrapper;
gr.regC = gr.registerControl;
gr.regCA = gr.registerControlArray;
gr.recCM = gr.setControlValue;

goradd.initialize();