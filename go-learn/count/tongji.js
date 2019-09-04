$(document).ready(function(){
    /**
     * 上报用户信息 以及访问数据到打点服务器
     */
    var url = ''; // 打点服务器的地址
    $.get(url,
        {
            "time"  : gettime(),
            "url"   : geturl(),
            "refer" : getrefer(),
            "ua"    : getuser_agent()
        })
})

/*
gettime(); // js获取当前时间
getip(); // js获取客户端ip
geturl(); // js获取客户端当前url
getrefer(); // js获取客户端当前页面的上级页面的url
getuser_agent(); // js获取客户端类型
*/
function gettime() {
    var nowDate = new Date();
    return nowDate.toLocaleDateString();
}

function getip() {
    return returnCitySN["cip"] + ',' + returnCitySN['cname'];
}

function getrefer() {
    return document.referrer;
}

function getcookie() {
    return document.cookie;
}

function getuser_agent() {
    return navigator.userAgent;
}