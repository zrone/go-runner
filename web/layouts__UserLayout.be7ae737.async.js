(self.webpackChunkant_design_pro=self.webpackChunkant_design_pro||[]).push([[81],{20698:function(q){q.exports={container:"container___3rwDa",lang:"lang___2ES0G",content:"content___3Paa8",top:"top___1W42Y",header:"header___1cl15",logo:"logo___29nS6",title:"title___3DxND",desc:"desc___2YLHe"}},2311:function(q,_,h){"use strict";h.r(_),h.d(_,{default:function(){return it}});var W=h(52663),y=h(67294),M=h(44721),p=h.n(M),st=h(97449),ft=h.n(st),G=h(78267),$=h.n(G),nt=h(23270),pt=h.n(nt);function A(){return(A=Object.assign||function(l){for(var e=1;e<arguments.length;e++){var r=arguments[e];for(var t in r)Object.prototype.hasOwnProperty.call(r,t)&&(l[t]=r[t])}return l}).apply(this,arguments)}function Z(l,e){l.prototype=Object.create(e.prototype),l.prototype.constructor=l,l.__proto__=e}function tt(l,e){if(l==null)return{};var r,t,n={},a=Object.keys(l);for(t=0;t<a.length;t++)e.indexOf(r=a[t])>=0||(n[r]=l[r]);return n}var s={BASE:"base",BODY:"body",HEAD:"head",HTML:"html",LINK:"link",META:"meta",NOSCRIPT:"noscript",SCRIPT:"script",STYLE:"style",TITLE:"title",FRAGMENT:"Symbol(react.fragment)"},H=Object.keys(s).map(function(l){return s[l]}),z={accesskey:"accessKey",charset:"charSet",class:"className",contenteditable:"contentEditable",contextmenu:"contextMenu","http-equiv":"httpEquiv",itemprop:"itemProp",tabindex:"tabIndex"},ct=Object.keys(z).reduce(function(l,e){return l[z[e]]=e,l},{}),w=function(e,r){for(var t=e.length-1;t>=0;t-=1){var n=e[t];if(Object.prototype.hasOwnProperty.call(n,r))return n[r]}return null},dt=function(e){var r=w(e,s.TITLE),t=w(e,"titleTemplate");if(Array.isArray(r)&&(r=r.join("")),t&&r)return t.replace(/%s/g,function(){return r});var n=w(e,"defaultTitle");return r||n||void 0},ht=function(e){return w(e,"onChangeClientState")||function(){}},X=function(e,r){return r.filter(function(t){return t[e]!==void 0}).map(function(t){return t[e]}).reduce(function(t,n){return A({},t,n)},{})},J=function(e,r){return r.filter(function(t){return t[s.BASE]!==void 0}).map(function(t){return t[s.BASE]}).reverse().reduce(function(t,n){if(!t.length)for(var a=Object.keys(n),o=0;o<a.length;o+=1){var i=a[o].toLowerCase();if(e.indexOf(i)!==-1&&n[i])return t.concat(n)}return t},[])},F=function(e,r,t){var n={};return t.filter(function(a){return!!Array.isArray(a[e])||(a[e]!==void 0&&console&&typeof console.warn=="function"&&console.warn("Helmet: "+e+' should be of type "Array". Instead found type "'+typeof a[e]+'"'),!1)}).map(function(a){return a[e]}).reverse().reduce(function(a,o){var i={};o.filter(function(T){for(var g,I=Object.keys(T),U=0;U<I.length;U+=1){var R=I[U],B=R.toLowerCase();r.indexOf(B)===-1||g==="rel"&&T[g].toLowerCase()==="canonical"||B==="rel"&&T[B].toLowerCase()==="stylesheet"||(g=B),r.indexOf(R)===-1||R!=="innerHTML"&&R!=="cssText"&&R!=="itemprop"||(g=R)}if(!g||!T[g])return!1;var et=T[g].toLowerCase();return n[g]||(n[g]={}),i[g]||(i[g]={}),!n[g][et]&&(i[g][et]=!0,!0)}).reverse().forEach(function(T){return a.push(T)});for(var u=Object.keys(i),f=0;f<u.length;f+=1){var d=u[f],O=A({},n[d],i[d]);n[d]=O}return a},[]).reverse()},lt=function(e){return Array.isArray(e)?e.join(""):e},mt=[s.NOSCRIPT,s.SCRIPT,s.STYLE],c=function(e,r){return r===void 0&&(r=!0),r===!1?String(e):String(e).replace(/&/g,"&amp;").replace(/</g,"&lt;").replace(/>/g,"&gt;").replace(/"/g,"&quot;").replace(/'/g,"&#x27;")},m=function(e){return Object.keys(e).reduce(function(r,t){var n=e[t]!==void 0?t+'="'+e[t]+'"':""+t;return r?r+" "+n:n},"")},E=function(e,r){return r===void 0&&(r={}),Object.keys(e).reduce(function(t,n){return t[z[n]||n]=e[n],t},r)},v=function(e,r,t){switch(e){case s.TITLE:return{toComponent:function(){return o=r.titleAttributes,(i={key:a=r.title})["data-rh"]=!0,u=E(o,i),[y.createElement(s.TITLE,u,a)];var a,o,i,u},toString:function(){return function(a,o,i,u){var f=m(i),d=lt(o);return f?"<"+a+' data-rh="true" '+f+">"+c(d,u)+"</"+a+">":"<"+a+' data-rh="true">'+c(d,u)+"</"+a+">"}(e,r.title,r.titleAttributes,t)}};case"bodyAttributes":case"htmlAttributes":return{toComponent:function(){return E(r)},toString:function(){return m(r)}};default:return{toComponent:function(){return function(a,o){return o.map(function(i,u){var f,d=((f={key:u})["data-rh"]=!0,f);return Object.keys(i).forEach(function(O){var T=z[O]||O;T==="innerHTML"||T==="cssText"?d.dangerouslySetInnerHTML={__html:i.innerHTML||i.cssText}:d[T]=i[O]}),y.createElement(a,d)})}(e,r)},toString:function(){return function(a,o,i){return o.reduce(function(u,f){var d=Object.keys(f).filter(function(g){return!(g==="innerHTML"||g==="cssText")}).reduce(function(g,I){var U=f[I]===void 0?I:I+'="'+c(f[I],i)+'"';return g?g+" "+U:U},""),O=f.innerHTML||f.cssText||"",T=mt.indexOf(a)===-1;return u+"<"+a+' data-rh="true" '+d+(T?"/>":">"+O+"</"+a+">")},"")}(e,r,t)}}}},L=function(e){var r=e.bodyAttributes,t=e.encode,n=e.htmlAttributes,a=e.linkTags,o=e.metaTags,i=e.noscriptTags,u=e.scriptTags,f=e.styleTags,d=e.title,O=d===void 0?"":d,T=e.titleAttributes;return{base:v(s.BASE,e.baseTag,t),bodyAttributes:v("bodyAttributes",r,t),htmlAttributes:v("htmlAttributes",n,t),link:v(s.LINK,a,t),meta:v(s.META,o,t),noscript:v(s.NOSCRIPT,i,t),script:v(s.SCRIPT,u,t),style:v(s.STYLE,f,t),title:v(s.TITLE,{title:O,titleAttributes:T},t)}},P=y.createContext({}),b=p().shape({setHelmet:p().func,helmetInstances:p().shape({get:p().func,add:p().func,remove:p().func})}),j=typeof document!="undefined",S=function(l){function e(r){var t;return(t=l.call(this,r)||this).instances=[],t.value={setHelmet:function(a){t.props.context.helmet=a},helmetInstances:{get:function(){return t.instances},add:function(a){t.instances.push(a)},remove:function(a){var o=t.instances.indexOf(a);t.instances.splice(o,1)}}},e.canUseDOM||(r.context.helmet=L({baseTag:[],bodyAttributes:{},encodeSpecialCharacters:!0,htmlAttributes:{},linkTags:[],metaTags:[],noscriptTags:[],scriptTags:[],styleTags:[],title:"",titleAttributes:{}})),t}return Z(e,l),e.prototype.render=function(){return y.createElement(P.Provider,{value:this.value},this.props.children)},e}(y.Component);S.canUseDOM=j,S.propTypes={context:p().shape({helmet:p().shape()}),children:p().node.isRequired},S.defaultProps={context:{}},S.displayName="HelmetProvider";var C=function(e,r){var t,n=document.head||document.querySelector(s.HEAD),a=n.querySelectorAll(e+"[data-rh]"),o=[].slice.call(a),i=[];return r&&r.length&&r.forEach(function(u){var f=document.createElement(e);for(var d in u)Object.prototype.hasOwnProperty.call(u,d)&&(d==="innerHTML"?f.innerHTML=u.innerHTML:d==="cssText"?f.styleSheet?f.styleSheet.cssText=u.cssText:f.appendChild(document.createTextNode(u.cssText)):f.setAttribute(d,u[d]===void 0?"":u[d]));f.setAttribute("data-rh","true"),o.some(function(O,T){return t=T,f.isEqualNode(O)})?o.splice(t,1):i.push(f)}),o.forEach(function(u){return u.parentNode.removeChild(u)}),i.forEach(function(u){return n.appendChild(u)}),{oldTags:o,newTags:i}},D=function(e,r){var t=document.getElementsByTagName(e)[0];if(t){for(var n=t.getAttribute("data-rh"),a=n?n.split(","):[],o=[].concat(a),i=Object.keys(r),u=0;u<i.length;u+=1){var f=i[u],d=r[f]||"";t.getAttribute(f)!==d&&t.setAttribute(f,d),a.indexOf(f)===-1&&a.push(f);var O=o.indexOf(f);O!==-1&&o.splice(O,1)}for(var T=o.length-1;T>=0;T-=1)t.removeAttribute(o[T]);a.length===o.length?t.removeAttribute("data-rh"):t.getAttribute("data-rh")!==i.join(",")&&t.setAttribute("data-rh",i.join(","))}},Q=function(e,r){var t=e.baseTag,n=e.htmlAttributes,a=e.linkTags,o=e.metaTags,i=e.noscriptTags,u=e.onChangeClientState,f=e.scriptTags,d=e.styleTags,O=e.title,T=e.titleAttributes;D(s.BODY,e.bodyAttributes),D(s.HTML,n),function(R,B){R!==void 0&&document.title!==R&&(document.title=lt(R)),D(s.TITLE,B)}(O,T);var g={baseTag:C(s.BASE,t),linkTags:C(s.LINK,a),metaTags:C(s.META,o),noscriptTags:C(s.NOSCRIPT,i),scriptTags:C(s.SCRIPT,f),styleTags:C(s.STYLE,d)},I={},U={};Object.keys(g).forEach(function(R){var B=g[R],et=B.newTags,vt=B.oldTags;et.length&&(I[R]=et),vt.length&&(U[R]=g[R].oldTags)}),r&&r(),u(e,I,U)},N=null,k=function(l){function e(){for(var t,n=arguments.length,a=new Array(n),o=0;o<n;o++)a[o]=arguments[o];return(t=l.call.apply(l,[this].concat(a))||this).rendered=!1,t}Z(e,l);var r=e.prototype;return r.shouldComponentUpdate=function(t){return!pt()(t,this.props)},r.componentDidUpdate=function(){this.emitChange()},r.componentWillUnmount=function(){this.props.context.helmetInstances.remove(this),this.emitChange()},r.emitChange=function(){var t,n,a=this.props.context,o=a.setHelmet,i=null,u=(t=a.helmetInstances.get().map(function(f){var d=A({},f.props);return delete d.context,d}),{baseTag:J(["href"],t),bodyAttributes:X("bodyAttributes",t),defer:w(t,"defer"),encode:w(t,"encodeSpecialCharacters"),htmlAttributes:X("htmlAttributes",t),linkTags:F(s.LINK,["rel","href"],t),metaTags:F(s.META,["name","charset","http-equiv","property","itemprop"],t),noscriptTags:F(s.NOSCRIPT,["innerHTML"],t),onChangeClientState:ht(t),scriptTags:F(s.SCRIPT,["src","innerHTML"],t),styleTags:F(s.STYLE,["cssText"],t),title:dt(t),titleAttributes:X("titleAttributes",t)});S.canUseDOM?(n=u,N&&cancelAnimationFrame(N),n.defer?N=requestAnimationFrame(function(){Q(n,function(){N=null})}):(Q(n),N=null)):L&&(i=L(u)),o(i)},r.init=function(){this.rendered||(this.rendered=!0,this.props.context.helmetInstances.add(this),this.emitChange())},r.render=function(){return this.init(),null},e}(y.Component);k.propTypes={context:b.isRequired},k.displayName="HelmetDispatcher";var K=function(l){function e(){return l.apply(this,arguments)||this}Z(e,l);var r=e.prototype;return r.shouldComponentUpdate=function(t){return!ft()(this.props,t)},r.mapNestedChildrenToProps=function(t,n){if(!n)return null;switch(t.type){case s.SCRIPT:case s.NOSCRIPT:return{innerHTML:n};case s.STYLE:return{cssText:n};default:throw new Error("<"+t.type+" /> elements are self-closing and can not contain children. Refer to our API for more information.")}},r.flattenArrayTypeChildren=function(t){var n,a=t.child,o=t.arrayTypeChildren;return A({},o,((n={})[a.type]=[].concat(o[a.type]||[],[A({},t.newChildProps,this.mapNestedChildrenToProps(a,t.nestedChildren))]),n))},r.mapObjectTypeChildren=function(t){var n,a,o=t.child,i=t.newProps,u=t.newChildProps,f=t.nestedChildren;switch(o.type){case s.TITLE:return A({},i,((n={})[o.type]=f,n.titleAttributes=A({},u),n));case s.BODY:return A({},i,{bodyAttributes:A({},u)});case s.HTML:return A({},i,{htmlAttributes:A({},u)});default:return A({},i,((a={})[o.type]=A({},u),a))}},r.mapArrayTypeChildrenToProps=function(t,n){var a=A({},n);return Object.keys(t).forEach(function(o){var i;a=A({},a,((i={})[o]=t[o],i))}),a},r.warnOnInvalidChildren=function(t,n){return $()(H.some(function(a){return t.type===a}),typeof t.type=="function"?"You may be attempting to nest <Helmet> components within each other, which is not allowed. Refer to our API for more information.":"Only elements types "+H.join(", ")+" are allowed. Helmet does not support rendering <"+t.type+"> elements. Refer to our API for more information."),$()(!n||typeof n=="string"||Array.isArray(n)&&!n.some(function(a){return typeof a!="string"}),"Helmet expects a string as a child of <"+t.type+">. Did you forget to wrap your children in braces? ( <"+t.type+">{``}</"+t.type+"> ) Refer to our API for more information."),!0},r.mapChildrenToProps=function(t,n){var a=this,o={};return y.Children.forEach(t,function(i){if(i&&i.props){var u=i.props,f=u.children,d=tt(u,["children"]),O=Object.keys(d).reduce(function(g,I){return g[ct[I]||I]=d[I],g},{}),T=i.type;switch(typeof T=="symbol"?T=T.toString():a.warnOnInvalidChildren(i,f),T){case s.FRAGMENT:n=a.mapChildrenToProps(f,n);break;case s.LINK:case s.META:case s.NOSCRIPT:case s.SCRIPT:case s.STYLE:o=a.flattenArrayTypeChildren({child:i,arrayTypeChildren:o,newChildProps:O,nestedChildren:f});break;default:n=a.mapObjectTypeChildren({child:i,newProps:n,newChildProps:O,nestedChildren:f})}}}),this.mapArrayTypeChildrenToProps(o,n)},r.render=function(){var t=this.props,n=t.children,a=A({},tt(t,["children"]));return n&&(a=this.mapChildrenToProps(n,a)),y.createElement(P.Consumer,null,function(o){return y.createElement(k,A({},a,{context:o}))})},e}(y.Component);K.propTypes={base:p().object,bodyAttributes:p().object,children:p().oneOfType([p().arrayOf(p().node),p().node]),defaultTitle:p().string,defer:p().bool,encodeSpecialCharacters:p().bool,htmlAttributes:p().object,link:p().arrayOf(p().object),meta:p().arrayOf(p().object),noscript:p().arrayOf(p().object),onChangeClientState:p().func,script:p().arrayOf(p().object),style:p().arrayOf(p().object),title:p().string,titleAttributes:p().object,titleTemplate:p().string},K.defaultProps={defer:!0,encodeSpecialCharacters:!0},K.displayName="Helmet";var Y=h(23715),ut=h(87748),rt=h(54200),V=h.n(rt),at=h(20698),x=h.n(at),ot=function(e){var r=e.children;return y.createElement(S,null,y.createElement(K,null,y.createElement("title",null,"Go Runner"),y.createElement("meta",{name:"description",content:"Go Runner"})),y.createElement("div",{className:x().container},y.createElement("div",{className:x().lang},y.createElement(Y.pD,null)),y.createElement("div",{className:x().content},y.createElement("div",{className:x().top},y.createElement("div",{className:x().header},y.createElement(ut.rU,{to:"/"},y.createElement("img",{alt:"logo",className:x().logo,src:V()}),y.createElement("span",{className:x().title},"Go Runner"))),y.createElement("div",{className:x().desc},y.createElement(Y._H,{id:"pages.layouts.userLayout.title",defaultMessage:"Go Runner \u662F\u6700\u7B80\u5355\u7684\u81EA\u52A8\u5316\u90E8\u7F72\u5DE5\u5177"}))),r)))},it=(0,Y.$j)(function(l){var e=l.settings;return(0,W.Z)({},e)})(ot)},87748:function(q,_,h){"use strict";h.d(_,{rU:function(){return w}});var W=h(2546),y=h(38279),M=h(67294),p=h(83233),st=h(44721),ft=h.n(st),G=h(3066),$=h(29345),nt=h(88945),pt=function(c){(0,y.Z)(m,c);function m(){for(var v,L=arguments.length,P=new Array(L),b=0;b<L;b++)P[b]=arguments[b];return v=c.call.apply(c,[this].concat(P))||this,v.history=(0,p.lX)(v.props),v}var E=m.prototype;return E.render=function(){return M.createElement(W.F0,{history:this.history,children:this.props.children})},m}(M.Component),A=function(c){(0,y.Z)(m,c);function m(){for(var v,L=arguments.length,P=new Array(L),b=0;b<L;b++)P[b]=arguments[b];return v=c.call.apply(c,[this].concat(P))||this,v.history=(0,p.q_)(v.props),v}var E=m.prototype;return E.render=function(){return M.createElement(W.F0,{history:this.history,children:this.props.children})},m}(M.Component),Z=function(m,E){return typeof m=="function"?m(E):m},tt=function(m,E){return typeof m=="string"?(0,p.ob)(m,null,null,E):m},s=function(m){return m},H=M.forwardRef;typeof H=="undefined"&&(H=s);function z(c){return!!(c.metaKey||c.altKey||c.ctrlKey||c.shiftKey)}var ct=H(function(c,m){var E=c.innerRef,v=c.navigate,L=c.onClick,P=(0,$.Z)(c,["innerRef","navigate","onClick"]),b=P.target,j=(0,G.Z)({},P,{onClick:function(C){try{L&&L(C)}catch(D){throw C.preventDefault(),D}!C.defaultPrevented&&C.button===0&&(!b||b==="_self")&&!z(C)&&(C.preventDefault(),v())}});return s!==H?j.ref=m||E:j.ref=E,M.createElement("a",j)}),w=H(function(c,m){var E=c.component,v=E===void 0?ct:E,L=c.replace,P=c.to,b=c.innerRef,j=(0,$.Z)(c,["component","replace","to","innerRef"]);return M.createElement(W.s6.Consumer,null,function(S){S||(0,nt.Z)(!1);var C=S.history,D=tt(Z(P,S.location),S.location),Q=D?C.createHref(D):"",N=(0,G.Z)({},j,{href:Q,navigate:function(){var K=Z(P,S.location),Y=L?C.replace:C.push;Y(K)}});return s!==H?N.ref=m||b:N.innerRef=b,M.createElement(v,N)})});if(!1)var dt,ht;var X=function(m){return m},J=M.forwardRef;typeof J=="undefined"&&(J=X);function F(){for(var c=arguments.length,m=new Array(c),E=0;E<c;E++)m[E]=arguments[E];return m.filter(function(v){return v}).join(" ")}var lt=J(function(c,m){var E=c["aria-current"],v=E===void 0?"page":E,L=c.activeClassName,P=L===void 0?"active":L,b=c.activeStyle,j=c.className,S=c.exact,C=c.isActive,D=c.location,Q=c.sensitive,N=c.strict,k=c.style,K=c.to,Y=c.innerRef,ut=(0,$.Z)(c,["aria-current","activeClassName","activeStyle","className","exact","isActive","location","sensitive","strict","style","to","innerRef"]);return M.createElement(W.s6.Consumer,null,function(rt){rt||(0,nt.Z)(!1);var V=D||rt.location,at=tt(Z(K,V),V),x=at.pathname,ot=x&&x.replace(/([.+*?=^!:${}()[\]|/\\])/g,"\\$1"),it=ot?(0,W.LX)(V.pathname,{path:ot,exact:S,sensitive:Q,strict:N}):null,l=!!(C?C(it,V):it),e=l?F(j,P):j,r=l?(0,G.Z)({},k,{},b):k,t=(0,G.Z)({"aria-current":l&&v||null,className:e,style:r,to:at},ut);return X!==J?t.ref=m||Y:t.innerRef=Y,M.createElement(w,t)})});if(!1)var mt},54200:function(q,_,h){q.exports=h.p+"static/logo.f0355d39.svg"}}]);
