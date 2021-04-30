//聊天记录页码
var ChatPage = 1
var PageStatus = 1


getChatLog()

//追加消息
//type :1 me 0 :you
// inputType 1 最新插入 0 末尾插入
function AddMsg(type , msg,inputType = 1){
    if( (typeof msg=='string')){
        var Msg = JSON.parse(msg)
    }else{
        var Msg =msg
    }
    let str = ""
    //群发或者私聊
    if (Msg.type === TypeSms || Msg.type === TypeSmsOne){
        var EmojiConfig = JSON.parse(Msg.head)
        if (type === 1){
            str =  ' <div class="message text-only">' +
                '          <div style="margin-right: 45px!important;" class="response">' +
                '            <div style="display: inline-block;">' +
                '              <p class="text">'+Msg.msg+'</p>' +
                '            </div>' +
                '          </div>' +
                '          <div  style="margin: 0;display: inline-block" class="main-content preview-head">' +
                '            <div style="transform: scale(0.33);background: none;max-height: 45px;position: relative;' +
                '    right: 20px;" class="emoji-preview my-head">' +
                '              <div class="emoji__wrapper">' +
                '                <div  class="emoji-face '+ EmojiConfig.skin+'">' +
                '                  <div class="hat '+EmojiConfig.hat+'"></div>' +
                '                  <div class="eyebrows">' +
                '                    <div class="eyebrow left '+EmojiConfig.eyebrow+'"></div>' +
                '                    <div class="eyebrow right '+EmojiConfig.eyebrow+'"></div>' +
                '                  </div>' +
                '                  <div class="eyes">' +
                '                    <div class="eye left '+EmojiConfig.eye+'"></div>' +
                '                    <div class="eye right '+EmojiConfig.eye+'"></div>' +
                '                  </div>' +
                '                  <div class="mouth '+EmojiConfig.mouth+'"></div>' +
                '                  <div class="face-extras '+EmojiConfig.faceExtras+'"></div>' +
                '                  <div class="item '+EmojiConfig.item+'"></div>' +
                '                </div>' +
                '              </div>' +
                '            </div>' +
                '          </div>' +
                '' +
                '        </div>'
        }
        else{
            str = '<div class="message">' +
                '          <div class="photo" style="height: 45px;position: relative;bottom:68px;background: none">' +
                '            <div class="main-content preview-head">' +
                '              <div style="transform: scale(0.33);background: none;" class="emoji-preview my-head">' +
                '                <div class="emoji__wrapper">' +
                '                  <div mail="'+Msg.mail+'" to_name="'+Msg.username+'" onclick="SendToUser(this)" class="emoji-face '+ EmojiConfig.skin+'">' +
                '                    <div class="hat '+EmojiConfig.hat+'"></div>' +
                '                    <div class="eyebrows">' +
                '                      <div class="eyebrow left '+EmojiConfig.eyebrow+'"></div>' +
                '                      <div class="eyebrow right '+EmojiConfig.eyebrow+'"></div>' +
                '                    </div>' +
                '                    <div class="eyes">' +
                '                      <div class="eye left '+EmojiConfig.eye+'"></div>' +
                '                      <div class="eye right '+EmojiConfig.eye+'"></div>' +
                '                    </div>' +
                '                    <div class="mouth '+EmojiConfig.mouth+'"></div>' +
                '                    <div class="face-extras '+EmojiConfig.faceExtras+'"></div>' +
                '                    <div class="item '+EmojiConfig.item+'"></div>' +
                '                  </div>' +
                '                </div>' +
                '              </div>' +
                '            </div>' +
                '          </div>' +
                '          <div style="display: inline-block">' +
                '            <span style="margin-left: 20px"> '+Msg.username+'  </span><br>' +
                '            <p class="text">'+Msg.msg+'</p>' +
                '          </div>' +
                '        </div>'
        }
    }else if(Msg.type === TypeRobot){
        str = '<div class="message">' +
            '            <div class="photo" style="background-image: url(/static/img/chatbot.jpg);">' +
            '            </div>' +
            ' <div style="display: inline-block">            ' +
            '   <span style="margin-left: 20px">'+ Msg.username+'</span>' +
            '   <br>' +
            '   <p class="text">'+ Msg.msg+'</p>          ' +
            ' </div>'+
            '</div>'

    } else{
        str ='<p class="time">'+Msg.msg+'</p>'
    }
    if(inputType === 1){
        $(".left_message").html(Msg.username+": "+Msg.msg)
        $(".messages-chat").append(str)
    }else{
        $(".messages-chat").prepend(str)
    }
}

//追加聊天时间
function AddTime(){
    var str ='<p class="time">17:30</p>'
    $(".messages-chat").append(str)
}

//私聊
function SendToUser(obj){
    return false
    var mail = $(obj).attr("mail")
    var to_name = $(obj).attr("to_name")
    $(".message-active").removeClass("message-active")
    var str =' <div class="discussion message-active">' +
        '        <div class="photo" style="">' +
        '          <div class="online"></div>' +
        '        </div>' +
        '        <div class="desc-contact">' +
        '          <p class="name">'+to_name+'</p>' +
        '          <p class="message"> </p>' +
        '        </div>' +
        '        <div class="timer">刚刚</div>' +
        '      </div>'
    $(".discussions").append(str)
}

$(document).ready(function(){
    $("#messages-chat").scroll(function () {
        var divHeight =$('.messages-chat').height();//div 高度
        var allHeight = $('.messages-chat')[0].scrollHeight; //总滚动长度
        var distance = $('.messages-chat')[0].scrollTop; //已经滚动距离
        if((divHeight+distance+50) >= allHeight){
            chatBottom();
        }
        if(distance === 0 ){
            if (PageStatus === 1){
                getChatLog(ChatPage)
            }
        }

    })
});


function AuthChat(){
    var divHeight =$('.messages-chat').height();//div 高度
    var allHeight = $('.messages-chat')[0].scrollHeight; //总滚动长度
    var distance = $('.messages-chat')[0].scrollTop; //已经滚动距离

    //在底部
    if((divHeight+distance+50) >= allHeight){
        chatBottom();
    }else{
        $(".newMsg").show(300)
    }
}


//聊天置底
function chatBottom(){
    $('.messages-chat').scrollTop( $('.messages-chat')[0].scrollHeight );
    $(".newMsg").hide()
}

function buildUrl(url){
    return url +"?key="+key+"&mail="+mail
}

$(document).on("click",".emoji-btn",function(e){
    e.stopPropagation();
    $('.intercom-composer-emoji-popover').toggleClass("active");
});
$(document).click(function (e) {
    if ($(e.target).attr('class') != '.intercom-composer-emoji-popover' && $(e.target).parents(".intercom-composer-emoji-popover").length == 0) {
        $(".intercom-composer-emoji-popover").removeClass("active");
    }
});

//搜索表情
$('.intercom-composer-popover-input').on('input', function() {
    var query = this.value;
    if(query != ""){
        $(".intercom-emoji-picker-emoji:not([title*='"+query+"'])").hide();
    }
    else{
        $(".intercom-emoji-picker-emoji").show();
    }
});

//键入表情
$(document).on("click",".intercom-emoji-picker-emoji",function(e){
    var oldVal = $("#msg").val()
    oldVal += $(this).html()
    $("#msg").val(oldVal);
});

//构建头像
function CreatHead(userinfo){
    var head = JSON.parse(userinfo.head)
    $(".my-head").empty()
    var str='<div class="emoji__wrapper">' +
        '              <div class="emoji-face '+head.skin+'">' +
        '                <div class="hat '+head.hat+'"></div>' +
        '                <div class="eyebrows">' +
        '                  <div class="eyebrow left '+head.eyebrow+'"></div>' +
        '                  <div class="eyebrow right '+head.eyebrow+'"></div>' +
        '                </div>' +
        '                <div class="eyes">' +
        '                  <div class="eye left '+head.eye+'"></div>' +
        '                  <div class="eye right '+head.eye+'"></div>' +
        '                </div>' +
        '                <div class="mouth '+head.mouth+'"></div>' +
        '                <div class="face-extras '+head.faceExtras+'"></div>' +
        '                <div class="item '+head.item+'"></div>' +
        '              </div>' +
        '            </div>'
    $(".my-head").html(str)
}


//新开页面
function Open(title,url,w,h,full) {
    if (title == null || title == '') {
        var title=false;
    };
    if (url == null || url == '') {
        var url= BaseRouter.Error404;
    };
    if (w == null || w == '') {
        var w=($(window).width()*0.8);
    };
    if (h == null || h == '') {
        var h=($(window).height() - 60);
    };
    var index = layer.open({
        type: 2,
        area: [w+'px', h +'px'],
        fix: false, //不固定
        maxmin: false,
        shadeClose: true,
        shade:0.4,
        title: title,
        content: url
    });
    if(full){
        layer.full(index);
    }
}


//退出登陆
function logout(is_wait_select =1){
    if (is_wait_select === 1){
        layer.confirm('确定退出登陆吗？', {
            btn: ['确定','取消'] //按钮
        }, function(){
            window.location.href='/view/logout'
        }, function(){
        });
    }else{
        window.location.href='/view/logout'
    }

}

//修改昵称
function rename(){
    layer.prompt({title: '请输入新的昵称(10字内)', formType: 3,value:UserInfo.username}, function(name, index){
        if(name.length >10){
            layer.msg("长度超过限制！",{icon:5})
            return false
        }
        layer.close(index);
        var load = layer.load()
        $.ajax({
            url:"/api/rename",
            dataType:"json",
            type:"get",
            data:{"newName":name},
            success:function (e) {
                layer.msg(e.msg,{icon:6})
                layer.close(load)
            },
            error:function (e){
                console.log(e)
                layer.msg("意外错误！",{icon:5})
                layer.close(load)
            }
        })

    });
}

function getChatLog(page=1,limit=5){
    setTimeout(function () {
        PageStatus = 0
        var load = layer.load()
        $.ajax({
            url:"/api/getChatLog",
            dataType:"json",
            type:"get",
            data:{"page":page,"limit":limit},
            success:function (e) {
                $.each(e,function (index,val) {
                    var msgType = 1
                    if (val.mail === UserInfo.mail){
                        msgType = 0
                    }
                    AddMsg(msgType,val,0)
                })
                ChatPage += 1
                PageStatus = 1
                layer.close(load)
            },
            error:function (e){
                console.log(e)
                PageStatus = 1
                layer.close(load)
            }
        })
    },1000)
}

