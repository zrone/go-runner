(self.webpackChunkant_design_pro=self.webpackChunkant_design_pro||[]).push([[541],{25061:function(me,M,t){"use strict";t.r(M),t.d(M,{default:function(){return de}});var pe=t(83385),w=t(55297),fe=t(44370),G=t(99957),he=t(41294),H=t(67739),y=t(51758),Z=t(40155),Ee=t(29068),f=t(65175),h=t(80153),K=t(94043),l=t.n(K),J=t(49101),e=t(67294),N=t(17044),Q=t(92220),R=t(72709),T=t(27398);function W(s){return V.apply(this,arguments)}function V(){return V=(0,h.Z)(l().mark(function s(u){return l().wrap(function(a){for(;;)switch(a.prev=a.next){case 0:return a.abrupt("return",(0,R.Z)("".concat(T.Z.host,"/task/proj/list"),{params:u}));case 1:case"end":return a.stop()}},s)})),V.apply(this,arguments)}function X(s){return $.apply(this,arguments)}function $(){return $=(0,h.Z)(l().mark(function s(u){return l().wrap(function(a){for(;;)switch(a.prev=a.next){case 0:return a.abrupt("return",(0,R.Z)("".concat(T.Z.host,"/task/proj/").concat(u),{method:"DELETE"}));case 1:case"end":return a.stop()}},s)})),$.apply(this,arguments)}function Y(s){return P.apply(this,arguments)}function P(){return P=(0,h.Z)(l().mark(function s(u){return l().wrap(function(a){for(;;)switch(a.prev=a.next){case 0:return a.abrupt("return",(0,R.Z)("".concat(T.Z.host,"/task/proj"),{method:"POST",data:(0,Z.Z)({},u)}));case 1:case"end":return a.stop()}},s)})),P.apply(this,arguments)}function q(s){return U.apply(this,arguments)}function U(){return U=(0,h.Z)(l().mark(function s(u){return l().wrap(function(a){for(;;)switch(a.prev=a.next){case 0:return a.abrupt("return",(0,R.Z)("".concat(T.Z.host,"/task/proj/").concat(u.symbol),{method:"PATCH",data:(0,Z.Z)({},u)}));case 1:case"end":return a.stop()}},s)})),U.apply(this,arguments)}var ve=t(41956),L=t(25017),j=t(28053),Fe=t(89167),O=t(23643),Ze=t(8241),z=t(58533),ye=t(55614),k=t(37455),be=t(59094),m=t(66032),Ce=t(76241),n=t(26308),x={labelCol:{span:5},wrapperCol:{span:18}},_={wrapperCol:{span:18,offset:5}},ee=function(u){var d=(0,e.useState)({name:"",secret:"",path:"",option:1,origin_tag:"",user:"",host:"",port:22,pwd:""}),a=(0,y.Z)(d,1),o=a[0],E=u.modalVisible,r=u.onCancel,D=u.onSubmit,b=n.Z.useForm(),C=(0,y.Z)(b,1),S=C[0],g=function(){var B=(0,h.Z)(l().mark(function v(){var A;return l().wrap(function(c){for(;;)switch(c.prev=c.next){case 0:return c.next=2,S.validateFields();case 2:A=c.sent,D((0,Z.Z)((0,Z.Z)({},o),A));case 4:case"end":return c.stop()}},v)}));return function(){return B.apply(this,arguments)}}();return e.createElement(L.Z,{bodyStyle:{padding:"32px 40px 10px"},destroyOnClose:!0,maskClosable:!1,title:"\u6DFB\u52A0\u9879\u76EE",visible:E,onCancel:function(){return r()},footer:null,width:800,centered:!0},e.createElement(n.Z,(0,j.Z)({},x,{form:S,colon:!1,initialValues:o,size:"middle"}),e.createElement(n.Z.Item,{name:"name",label:"\u9879\u76EE\u540D\u79F0",hasFeedback:!0,rules:[{required:!0,message:"\u9879\u76EE\u540D\u79F0\u4E0D\u80FD\u4E3A\u7A7A\uFF01"}]},e.createElement(m.Z,{placeholder:"\u8BF7\u8F93\u5165"})),e.createElement(n.Z.Item,{name:"option",label:"\u90E8\u7F72\u7C7B\u578B",hasFeedback:!0},e.createElement(k.ZP.Group,null,e.createElement(k.ZP,{value:1},"\u81EA\u52A8\u5316\u90E8\u7F72"),e.createElement(k.ZP,{value:2},"\u4E0A\u7EBF\u53D1\u5E03"))),e.createElement(n.Z.Item,{name:"origin_tag",label:"\u53D1\u5E03\u5206\u652F",hasFeedback:!0,tooltip:"\u4E0A\u7EBF\u53D1\u5E03\u65F6\u5FC5\u987B\uFF0C\u586B\u5199branch\u6216tag"},e.createElement(m.Z,{placeholder:"\u8BF7\u8F93\u5165"})),e.createElement(n.Z.Item,{name:"secret",label:"\u79D8\u94A5",hasFeedback:!0,tooltip:"\u81EA\u52A8\u5316\u90E8\u7F72\u5FC5\u987B\u914D\u7F6E\u6821\u9A8C\u79D8\u94A5"},e.createElement(m.Z,{placeholder:"\u8BF7\u8F93\u5165"})),e.createElement(n.Z.Item,{name:"path",label:"\u90E8\u7F72\u8DEF\u5F84",hasFeedback:!0,rules:[{required:!0,message:"\u90E8\u7F72\u8DEF\u5F84\u4E0D\u80FD\u4E3A\u7A7A\uFF01"}]},e.createElement(m.Z,{placeholder:"\u8BF7\u8F93\u5165"})),e.createElement(n.Z.Item,{name:"host",label:"\u4E3B\u673A",hasFeedback:!0,rules:[{required:!0,message:"\u4E3B\u673A\u4E0D\u80FD\u4E3A\u7A7A\uFF01"}]},e.createElement(m.Z,{placeholder:"\u8BF7\u8F93\u5165"})),e.createElement(n.Z.Item,{name:"user",label:"\u7528\u6237\u540D",hasFeedback:!0,rules:[{required:!0,message:"\u7528\u6237\u540D\u4E0D\u80FD\u4E3A\u7A7A\uFF01"}]},e.createElement(m.Z,{placeholder:"\u8BF7\u8F93\u5165"})),e.createElement(n.Z.Item,{name:"port",label:"\u7AEF\u53E3",rules:[{required:!0,message:"\u7AEF\u53E3\u4E0D\u80FD\u4E3A\u7A7A!"}]},e.createElement(z.Z,{min:0,placeholder:"\u8BF7\u8F93\u5165"})),e.createElement(n.Z.Item,{name:"pwd",label:"\u767B\u5F55\u5BC6\u7801",hasFeedback:!0,rules:[{required:!0,message:"\u767B\u5F55\u5BC6\u7801\u4E0D\u80FD\u4E3A\u7A7A\uFF01"}]},e.createElement(m.Z,{placeholder:"\u8BF7\u8F93\u5165"})),e.createElement(n.Z.Item,_,e.createElement(O.Z,null,e.createElement(w.Z,{type:"primary",onClick:function(){return g()}},"\u63D0\u4EA4"),e.createElement(w.Z,{onClick:r},"\u53D6\u6D88")))))},re=ee,ae={labelCol:{span:5},wrapperCol:{span:18}},te={wrapperCol:{span:18,offset:5}},ue=function(u){var d=u.onSubmit,a=u.onCancel,o=u.updateModalVisible,E=u.values,r=(0,e.useState)(E),D=(0,y.Z)(r,1),b=D[0],C=n.Z.useForm(),S=(0,y.Z)(C,1),g=S[0],B=function(){var v=(0,h.Z)(l().mark(function A(){var p;return l().wrap(function(i){for(;;)switch(i.prev=i.next){case 0:return i.next=2,g.validateFields();case 2:p=i.sent,d((0,Z.Z)((0,Z.Z)({},b),p));case 4:case"end":return i.stop()}},A)}));return function(){return v.apply(this,arguments)}}();return e.createElement(L.Z,{bodyStyle:{padding:"32px 40px 10px"},destroyOnClose:!0,maskClosable:!1,title:"\u7F16\u8F91\u9879\u76EE",visible:o,onCancel:function(){return a()},footer:null,width:800,centered:!0},e.createElement(n.Z,(0,j.Z)({},ae,{form:g,colon:!1,initialValues:b,size:"middle"}),e.createElement(n.Z.Item,{name:"symbol",label:"\u9879\u76EE\u6807\u8BC6",hasFeedback:!0,rules:[{required:!0,message:"\u9879\u76EE\u6807\u8BC6\u4E0D\u80FD\u4E3A\u7A7A\uFF01"}]},e.createElement(m.Z,{disabled:!0,placeholder:"\u8BF7\u8F93\u5165"})),e.createElement(n.Z.Item,{name:"name",label:"\u9879\u76EE\u540D\u79F0",hasFeedback:!0,rules:[{required:!0,message:"\u9879\u76EE\u540D\u79F0\u4E0D\u80FD\u4E3A\u7A7A\uFF01"}]},e.createElement(m.Z,{placeholder:"\u8BF7\u8F93\u5165"})),e.createElement(n.Z.Item,{name:"option",label:"\u90E8\u7F72\u7C7B\u578B",hasFeedback:!0},e.createElement(k.ZP.Group,null,e.createElement(k.ZP,{value:1},"\u81EA\u52A8\u5316\u90E8\u7F72"),e.createElement(k.ZP,{value:2},"\u4E0A\u7EBF\u53D1\u5E03"))),e.createElement(n.Z.Item,{name:"origin_tag",label:"\u53D1\u5E03\u5206\u652F",hasFeedback:!0,tooltip:"\u4E0A\u7EBF\u53D1\u5E03\u65F6\u5FC5\u987B\uFF0C\u586B\u5199branch\u6216tag"},e.createElement(m.Z,{placeholder:"\u8BF7\u8F93\u5165"})),e.createElement(n.Z.Item,{name:"secret",label:"\u79D8\u94A5",hasFeedback:!0,tooltip:"\u81EA\u52A8\u5316\u90E8\u7F72\u5FC5\u987B\u914D\u7F6E\u6821\u9A8C\u79D8\u94A5"},e.createElement(m.Z,{placeholder:"\u8BF7\u8F93\u5165"})),e.createElement(n.Z.Item,{name:"path",label:"\u90E8\u7F72\u8DEF\u5F84",hasFeedback:!0,rules:[{required:!0,message:"\u90E8\u7F72\u8DEF\u5F84\u4E0D\u80FD\u4E3A\u7A7A\uFF01"}]},e.createElement(m.Z,{placeholder:"\u8BF7\u8F93\u5165"})),e.createElement(n.Z.Item,{name:"host",label:"\u4E3B\u673A",hasFeedback:!0,rules:[{required:!0,message:"\u4E3B\u673A\u4E0D\u80FD\u4E3A\u7A7A\uFF01"}]},e.createElement(m.Z,{placeholder:"\u8BF7\u8F93\u5165"})),e.createElement(n.Z.Item,{name:"user",label:"\u7528\u6237\u540D",hasFeedback:!0,rules:[{required:!0,message:"\u7528\u6237\u540D\u4E0D\u80FD\u4E3A\u7A7A\uFF01"}]},e.createElement(m.Z,{placeholder:"\u8BF7\u8F93\u5165"})),e.createElement(n.Z.Item,{name:"port",label:"\u7AEF\u53E3",rules:[{required:!0,message:"\u7AEF\u53E3\u4E0D\u80FD\u4E3A\u7A7A!"}]},e.createElement(z.Z,{min:0,placeholder:"\u8BF7\u8F93\u5165"})),e.createElement(n.Z.Item,{name:"pwd",label:"\u767B\u5F55\u5BC6\u7801",hasFeedback:!0},e.createElement(m.Z,{placeholder:"\u8BF7\u8F93\u5165"})),e.createElement(n.Z.Item,te,e.createElement(O.Z,null,e.createElement(w.Z,{type:"primary",onClick:function(){return B()}},"\u63D0\u4EA4"),e.createElement(w.Z,{onClick:function(){return a(!1,E)}},"\u53D6\u6D88")))))},ne=ue,le=function(){var s=(0,h.Z)(l().mark(function u(d){var a,o;return l().wrap(function(r){for(;;)switch(r.prev=r.next){case 0:return a=f.default.loading({content:"\u63D0\u4EA4\u4E2D...",key:"handleAdd"}),r.prev=1,r.next=4,Y((0,Z.Z)({},d));case 4:if(o=r.sent,o.code!=400){r.next=8;break}return f.default.error({content:o.message,key:"handleAdd",duration:2}),r.abrupt("return",!1);case 8:return f.default.success({content:"\u521B\u5EFA\u6210\u529F!",key:"handleAdd",duration:2}),r.abrupt("return",!0);case 12:return r.prev=12,r.t0=r.catch(1),f.default.error({content:"\u521B\u5EFA\u5931\u8D25!",key:"handleAdd",duration:2}),r.abrupt("return",!1);case 16:return r.prev=16,setTimeout(function(){a()},2e3),r.finish(16);case 19:case"end":return r.stop()}},u,null,[[1,12,16,19]])}));return function(d){return s.apply(this,arguments)}}(),se=function(){var s=(0,h.Z)(l().mark(function u(d){var a,o;return l().wrap(function(r){for(;;)switch(r.prev=r.next){case 0:return a=f.default.loading({content:"\u63D0\u4EA4\u4E2D...",key:"handleAdd"}),r.prev=1,r.next=4,q(d);case 4:if(o=r.sent,o.code!=400){r.next=8;break}return f.default.error({content:o.message,key:"handleAdd",duration:2}),r.abrupt("return",!1);case 8:return f.default.success({content:"\u66F4\u65B0\u6210\u529F!",key:"handleAdd",duration:2}),r.abrupt("return",!0);case 12:return r.prev=12,r.t0=r.catch(1),f.default.error({content:"\u66F4\u65B0\u5931\u8D25!",key:"handleAdd",duration:2}),r.abrupt("return",!1);case 16:return r.prev=16,setTimeout(function(){a()},2e3),r.finish(16);case 19:case"end":return r.stop()}},u,null,[[1,12,16,19]])}));return function(d){return s.apply(this,arguments)}}(),oe=function(){var s=(0,h.Z)(l().mark(function u(d){var a,o;return l().wrap(function(r){for(;;)switch(r.prev=r.next){case 0:return a=f.default.loading({content:"\u63D0\u4EA4\u4E2D...",key:"handleRemove"}),r.prev=1,r.next=4,X(d);case 4:if(o=r.sent,o.code!=400){r.next=8;break}return f.default.error({content:o.message,key:"handleRemove",duration:2}),r.abrupt("return",!1);case 8:return f.default.success({content:"\u5220\u9664\u6210\u529F!",key:"handleRemove",duration:2}),r.abrupt("return",!0);case 12:return r.prev=12,r.t0=r.catch(1),f.default.error({content:"\u5220\u9664\u5931\u8D25!",key:"handleRemove",duration:2}),r.abrupt("return",!1);case 16:return r.prev=16,setTimeout(function(){a()},2e3),r.finish(16);case 19:case"end":return r.stop()}},u,null,[[1,12,16,19]])}));return function(d){return s.apply(this,arguments)}}(),ie=function(){var u=(0,e.useRef)(),d=(0,e.useState)(!1),a=(0,y.Z)(d,2),o=a[0],E=a[1],r=(0,e.useState)(!1),D=(0,y.Z)(r,2),b=D[0],C=D[1],S=(0,e.useState)(),g=(0,y.Z)(S,2),B=g[0],v=g[1],A=[{title:"\u9879\u76EE\u540D\u79F0 / \u6807\u8BC6",dataIndex:"name",fixed:"left",ellipsis:!0,width:180,render:function(c,i){return e.createElement(e.Fragment,null,e.createElement("span",null,i.name,e.createElement("br",null)),e.createElement("span",null,i.symbol))}},{title:"\u9879\u76EE\u7C7B\u578B",dataIndex:"option",width:100,render:function(c,i){return i.option==1?"\u81EA\u52A8\u5316\u90E8\u7F72":"\u4E0A\u7EBF\u53D1\u5E03"}},{title:"\u79D8\u94A5",dataIndex:"secret",ellipsis:!0,copyable:!0,width:220,hideInTable:!0},{title:"\u90E8\u7F72\u76EE\u5F55",dataIndex:"path",ellipsis:!0,width:150},{title:"\u4E3B\u673A\u5730\u5740",dataIndex:"host",hideInTable:!0},{title:"\u7528\u6237\u540D",dataIndex:"ssh",hideInTable:!0},{title:"\u7AEF\u53E3",dataIndex:"port",hideInTable:!0},{title:"\u94FE\u63A5\u5BC6\u7801",dataIndex:"pwd",hideInTable:!0},{title:"\u64CD\u4F5C",dataIndex:"option",valueType:"option",width:120,fixed:"right",render:function(c,i){return e.createElement(e.Fragment,null,e.createElement("a",{onClick:function(){v(i),C(!0)}},"\u914D\u7F6E"),e.createElement(H.Z,{type:"vertical"}),e.createElement(G.Z,{title:"\u786E\u8BA4\u8981\u5220\u9664\u9879\u76EE\u914D\u7F6E\u5417?",onConfirm:function(){oe(i.symbol),u.current&&u.current.reload()},okText:"\u786E\u8BA4",cancelText:"\u53D6\u6D88"},e.createElement("a",null,"\u5220\u9664")))}}];return e.createElement(N.Z,null,e.createElement(Q.ZP,{headerTitle:"\u9879\u76EE\u5217\u8868",size:"small",actionRef:u,rowKey:"symbol",search:!1,toolBarRender:function(){return[e.createElement(w.Z,{type:"primary",key:"primary",onClick:function(){E(!0)}},e.createElement(J.Z,null)," \u6DFB\u52A0")]},request:function(c,i){return W((0,Z.Z)((0,Z.Z)({},c),{},{sorter:i}))},columns:A}),o?e.createElement(re,{onCancel:function(){return E(!1)},modalVisible:o,onSubmit:function(){var p=(0,h.Z)(l().mark(function c(i){var I;return l().wrap(function(F){for(;;)switch(F.prev=F.next){case 0:return F.next=2,le(i);case 2:I=F.sent,I&&(E(!1),u.current&&u.current.reload());case 4:case"end":return F.stop()}},c)}));return function(c){return p.apply(this,arguments)}}()}):null,b?e.createElement(ne,{values:B,onSubmit:function(){var p=(0,h.Z)(l().mark(function c(i){var I;return l().wrap(function(F){for(;;)switch(F.prev=F.next){case 0:return F.next=2,se(i);case 2:I=F.sent,I&&(C(!1),v({symbol:"",name:"",secret:"",path:"",option:1,origin_tag:"",user:"",host:"",port:22,pwd:""}),u.current&&u.current.reload());case 4:case"end":return F.stop()}},c)}));return function(c){return p.apply(this,arguments)}}(),onCancel:function(){C(!1),v({symbol:"",name:"",secret:"",path:"",option:1,origin_tag:"",user:"",host:"",port:22,pwd:""})},updateModalVisible:b}):null)},de=ie}}]);