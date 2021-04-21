



//追加消息
//type :1 me 0 :you
function AddMsg(type , msg){
    var Msg = JSON.parse(msg)
    console.log(Msg)
    let str = ""
    //群发或者私聊
    if (Msg.type === TypeSms || Msg.type === TypeSmsOne){
        if (type === 1){
            str = '<div class="message text-only">\n' +
                '          <div class="response">\n' +
                '            <div style="display: inline-block;">\n' +
                '              <p class="text"> '+Msg.msg+'</p>\n' +
                '            </div>\n' +
                '            <div  class="photo" style="background-image: url(http://e0.365dm.com/16/08/16-9/20/theirry-henry-sky-sports-pundit_3766131.jpg?20161212144602);display: inline-block">\n' +
                '            </div>\n' +
                '          </div>\n' +
                '        </div>'
        }else{
            str =  '<div class="message">\n' +
                '          <div class="photo" style="background-image: url(http://e0.365dm.com/16/08/16-9/20/theirry-henry-sky-sports-pundit_3766131.jpg?20161212144602);">\n' +
                '          </div>\n' +
                '          <div style="display: inline-block;margin-top: 30px">\n' +
                '            <span style="margin-left: 20px"> '+Msg.username+'  </span><br>\n' +
                '            <p class="text">'+Msg.msg+'</p>\n' +
                '          </div>\n' +
                '        </div>'
        }
    }else{
        str ='<p class="time">'+Msg.msg+'</p>'
    }
    $(".messages-chat").append(str)
}

//追加聊天时间
function AddTime(){
    var str ='<p class="time">17:30</p>'
    $(".messages-chat").append(str)
}

function buildUrl(url){
    return url +"?key="+key+"&mail="+mail
}