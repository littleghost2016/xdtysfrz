// 全局参数，是否需要验证码
var needCaptcha = false;

(function () {
    if (getObj("errorMsg")) {
        showTips(getObj("errorMsg").innerHTML);
    }

    // 验证码刷新事件
    getObj("captchaImg").onclick = function () {
        getObj("captchaImg").src = "captcha.html?ts=" + new Date().getMilliseconds();
    };
    //滑块验证码：不需要用户名离焦
    var isSliderCaptcha=$("#isSliderCaptcha").val();
    if(!isSliderCaptcha) {
        //  用户名输入框修改事件，判断是否需要验证码
        getObj("mobileUsername").onblur = function () {
            getCaptcha();
        };

        // 密码输入框修改事件，判断是否需要验证码
        getObj("mobilePassword").onblur = function () {
            getCaptcha();
        };
    }
    //滑块验证码：这里是随机验证码
    if(!isSliderCaptcha) {
        // 帐号登陆提交事件
        getObj("casLoginForm").onsubmit = function doLogin() {
            var username = getObj("mobileUsername");
            var password = getObj("mobilePassword");
            var captchaResponse = getObj("captchaResponse");

            if (!checkRequired(username, getObj("usernameEmpty").innerText)) {
                username.focus();
                return false;
            }
            if (!checkRequired(password, getObj("passwordEmpty").innerText)) {
                password.focus();
                return false;
            }
            if (needCaptcha) {
                if (!checkRequired(captchaResponse, getObj("captchaEmpty").innerText)) {
                    captchaResponse.focus();
                    return false;
                }
            }
            _etd(password.value);
            getObj("load").disabled = true;
            getObj("loading").style.width = "1em";
            //        return true;
        }
    }
})();
//滑块：帐号登陆表单
function submitLoginForm(e) {
    e.preventDefault();
    var username = getObj("mobileUsername");
    var password = getObj("mobilePassword");
    if (!checkRequired(username, getObj("usernameEmpty").innerText)) {
        username.focus();
        return false;
    }
    if (!checkRequired(password, getObj("passwordEmpty").innerText)) {
        password.focus();
        return false;
    }
    _etd(password.value);
    var isSliderCaptcha=$("#isSliderCaptcha").val();
    reValidateDeal(isSliderCaptcha);
}

function reValidateDeal(isSliderCaptcha) {
    var username = getObj("mobileUsername").value;
    if (username != "") {
        ajax({
            url: "needCaptcha.html",
            data: {username: username, pwdEncrypt2: "pwdEncryptSalt"},
            dataType: "json",
            success: function (data) {
                if (data.indexOf("::::") > -1) {
                    var pwdEncryptArr = data.split("::::");
                    try {
                        pwdDefaultEncryptSalt = pwdEncryptArr[1];
                    } catch (e) {
                    }
                }
                if (data.indexOf("true") > -1) {
                    //如果是滑块验证码
                    if (isSliderCaptcha) {
                        $("#captcha-id").show();
                        createSliderCaptcha();
                        //区分动态码跟账号密码登录
                        $("#sliderCaptchaDynamicCode").val("");
                    }
                } else {
                    $("#captcha-id").empty();
                    //登录不需要滑块验证码，直接登录
                    var casLoginForm = getObj("casLoginForm");
                    casLoginForm.submit();
                }
            }
        });
    }
}

// 统一校验必填和展示错误信息的方法
function checkRequired(obj, msg) {
    if (obj.value == "") {
        showTips(msg);
        return false;
    } else {
        return true;
    }
}

function showCaptcha() {
    needCaptcha = true;
    getObj("cpatchaDiv").style.display = "";
}

function hideCaptcha() {
    needCaptcha = false;
    getObj("cpatchaDiv").style.display = "none";
}

function _etd(_p0) {
    try {
        var _p2 = encryptAES(_p0, pwdDefaultEncryptSalt);
        getObj("mobilePasswordEncrypt").value = _p2;
    } catch (e) {
        getObj("mobilePasswordEncrypt").value = _p0;
    }
}

function getCaptcha() {
    var username = getObj("mobileUsername").value;
    if (username != "") {
        ajax({
            url: "needCaptcha.html",
            data: {username: username, pwdEncrypt2: "pwdEncryptSalt"},
            dataType: "json",
            success: function (data) {
                if (data.indexOf("::::") > -1) {
                    var pwdEncryptArr = data.split("::::");
                    try {
                        pwdDefaultEncryptSalt = pwdEncryptArr[1];
                    } catch (e) {
                    }
                }
                if (data.indexOf("true") > -1) {
                    showCaptcha();
                } else {
                    hideCaptcha();
                }
            }
        });
    }
}