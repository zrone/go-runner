(self.webpackChunkant_design_pro=self.webpackChunkant_design_pro||[]).push([[694],{69524:function(x,i,e){"use strict";e.r(i);var c=e(81766),R=e(7935),l=e(69885),m=e(87379),D=e(89167),h=e(23643),O=e(44370),U=e(99957),M=e(83385),P=e(55297),v=e(1898),r=e(29024),o=e(51758),T=e(9950),t=e(20301),Ee=e(29068),f=e(65175),j=e(80153),z=e(94043),p=e.n(z),n=e(67294),$=e(17044),L=e(18701),me=e.n(L),I=e(9688),de=e.n(I),W=e(99625),ie=e.n(W),F=e(27398),N=e(99963),B=e.n(N),G=e(25747),H=e.n(G),J=e(17993),X=function(){var K=(0,j.Z)(p().mark(function C(s){var y,u;return p().wrap(function(_){for(;;)switch(_.prev=_.next){case 0:return y=f.default.loading({content:"\u63D0\u4EA4\u4E2D...",key:"handleRetry"}),_.prev=1,_.next=4,(0,J.XD)(s);case 4:u=_.sent,u.code==400?f.default.error({content:u.message,key:"handleRetry",duration:2}):f.default.success({content:u.message,key:"handleRetry",duration:2}),_.next=11;break;case 8:_.prev=8,_.t0=_.catch(1),f.default.error({content:"\u91CD\u8BD5\u5931\u8D25!",key:"handleRetry",duration:2});case 11:return _.prev=11,setTimeout(function(){y()},2e3),_.finish(11);case 14:case"end":return _.stop()}},C,null,[[1,8,11,14]])}));return function(s){return K.apply(this,arguments)}}(),k=function(C){var s=C.location.query,y=s.symbol,u=s.pipline,Z=s.version,_=s.committer,ce=s.create_at,Q=s.end_at,V=t.Z.Header,Y=t.Z.Content,w=H()(Q*1e3).startOf("second").fromNow(),g=new L.Terminal({fontSize:14,disableStdin:!1,cursorStyle:"bar"}),b=new W.FitAddon;g.loadAddon(new I.WebLinksAddon),g.loadAddon(b);var q=(0,n.useState)(b),ee=(0,o.Z)(q,1),te=ee[0],ne=(0,n.useState)(g),_e=(0,o.Z)(ne,1),d=_e[0],ae=function(){B()(".tml .ant-card-body").height(B()(".rightMenu").height()-64+"px"),B()(".tml #terminal").height(B()(".rightMenu").height()-64+"px");var a=new WebSocket(F.Z.wss);se(a),re(a),oe(a),le(a)},re=function(a){a.onopen=function(){d.open(document.getElementById("terminal")),te.fit(),a.send(JSON.stringify({UUID:y}))}},se=function(a){a.onclose=function(){}},oe=function(a){a.onerror=function(){console.log("socket \u94FE\u63A5\u5931\u8D25")}},le=function(a){a.onmessage=function(ue){var S=JSON.parse(ue.data),A=S.data.msg.trim()+`\r
`;switch(S.data.level){case"info":d.write("[1;37m"+A+"[0m");break;case"error":d.write("[1;31m"+A+"[0m");break;case"debug":d.write("[1;32m"+A+"[0m");break;default:d.write("[1;35m"+A+"[0m");break}}};return(0,n.useEffect)(function(){ae()},[]),n.createElement($.Z,{className:"tml"},n.createElement(R.Z,null,n.createElement(t.Z,null,n.createElement(V,{className:"t-header"},n.createElement(m.Z,null,n.createElement(r.Z,{span:8},n.createElement("b",null,"#",Z),"\xA0",n.createElement("span",null,"triggered ",w," by"),"\xA0",_),n.createElement(r.Z,{span:8,offset:8,style:{textAlign:"right"}},n.createElement(h.Z,null,n.createElement(U.Z,{title:"\u60A8\u786E\u8BA4\u8981\u91CD\u8BD5\u6B64\u90E8\u7F72\u4EFB\u52A1\u5417?",onConfirm:function(){X(u)},okText:"\u91CD\u8BD5",cancelText:"\u53D6\u6D88"},n.createElement(P.Z,{danger:!0},"\u91CD\u8BD5")),n.createElement(P.Z,{onClick:function(){history.go(-1)},type:"primary"},"\u8FD4\u56DE"))))),n.createElement(Y,null,n.createElement("div",{id:"terminal"})))))};i.default=k},17993:function(x,i,e){"use strict";e.d(i,{oF:function(){return h},c_:function(){return U},XD:function(){return P}});var c=e(80153),R=e(94043),l=e.n(R),m=e(72709),D=e(27398);function h(r){return O.apply(this,arguments)}function O(){return O=(0,c.Z)(l().mark(function r(o){return l().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,m.Z)("".concat(D.Z.host,"/task/console/list"),{params:o}));case 1:case"end":return t.stop()}},r)})),O.apply(this,arguments)}function U(r){return M.apply(this,arguments)}function M(){return M=(0,c.Z)(l().mark(function r(o){return l().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,m.Z)("".concat(D.Z.host,"/task/release/list"),{params:o}));case 1:case"end":return t.stop()}},r)})),M.apply(this,arguments)}function P(r){return v.apply(this,arguments)}function v(){return v=(0,c.Z)(l().mark(function r(o){return l().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,m.Z)("".concat(D.Z.host,"/task/retry/").concat(o),{method:"POST",data:{}}));case 1:case"end":return t.stop()}},r)})),v.apply(this,arguments)}}}]);
