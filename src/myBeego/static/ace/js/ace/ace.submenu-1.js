ace.submenu={show:function(sub,duration){var $=window.jQuery;var $sub=$(sub);var event;$sub.trigger(event=$.Event("show.ace.submenu"));if(event.isDefaultPrevented()){return false}$sub.css({height:0,overflow:"hidden",display:"block"}).removeClass("nav-hide").addClass("nav-show").parent().addClass("open");if(duration>0){$sub.css({height:sub.scrollHeight,"transition-property":"height","transition-duration":(duration/1000)+"s"})}var complete=function(ev,trigger){ev&&ev.stopPropagation();$sub.css({"transition-property":"","transition-duration":"",overflow:"",height:""});if(ace.vars["transition"]){$sub.off(".trans")}if(trigger!==false){$sub.trigger($.Event("shown.ace.submenu"))}};if(duration>0&&ace.vars["transition"]){$sub.one("transitionend.trans webkitTransitionEnd.trans mozTransitionEnd.trans oTransitionEnd.trans",complete)}else{complete()}if(ace.vars["android"]){setTimeout(function(){complete(null,false)},duration+20)}return true},hide:function(sub,duration){var $=window.jQuery;var $sub=$(sub);var event;$sub.trigger(event=$.Event("hide.ace.submenu"));if(event.isDefaultPrevented()){return false}$sub.css({height:sub.scrollHeight,overflow:"hidden"}).parent().removeClass("open");sub.offsetHeight;if(duration>0){$sub.css({"height":0,"transition-property":"height","transition-duration":(duration/1000)+"s"})}var complete=function(ev,trigger){ev&&ev.stopPropagation();$sub.css({display:"none",overflow:"",height:"","transition-property":"","transition-duration":""}).removeClass("nav-show").addClass("nav-hide");if(ace.vars["transition"]){$sub.off(".trans")}if(trigger!==false){$sub.trigger($.Event("hidden.ace.submenu"))}};if(duration>0&&ace.vars["transition"]){$sub.one("transitionend.trans webkitTransitionEnd.trans mozTransitionEnd.trans oTransitionEnd.trans",complete)}else{complete()}if(ace.vars["android"]){setTimeout(function(){complete(null,false)},duration+20)}return true},toggle:function(element,duration){if(element.scrollHeight==0){if(ace.submenu.show(element,duration)){return 1}}else{if(ace.submenu.hide(element,duration)){return -1}}return 0}};