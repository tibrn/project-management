(window.jsonpFunction=window.jsonpFunction||[]).push([["layouts"],{"25b3":function(t,e,i){"use strict";i.r(e);var n,o=i("d225"),s=i("b0b4"),a=i("308d"),r=i("6bb5"),c=i("4e2b"),l=i("60a3"),h=Object(l.a)({name:"AuthLayout"})(n=function(t){function e(){return Object(o.a)(this,e),Object(a.a)(this,Object(r.a)(e).apply(this,arguments))}return Object(c.a)(e,t),Object(s.a)(e,[{key:"created",value:function(){console.log("Auth-Layout")}}]),e}(l.b))||n,u=i("2877"),d=i("6544"),p=i.n(d),f=i("a523"),v=i("549c"),b=Object(u.a)(h,function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"Auth"},[e("v-content",[e("v-container",{attrs:{fluid:""}},[e("router-view")],1)],1)],1)},[],!1,null,null,null);e.default=b.exports,p()(b,{VContainer:f.a,VContent:v.a})},"549c":function(t,e,i){"use strict";i("f134");var n=i("b57a");e.a={name:"v-content",mixins:[n.a],props:{tag:{type:String,default:"main"}},computed:{styles:function(){var t=this.$vuetify.application,e=t.bar;return{paddingTop:t.top+e+"px",paddingRight:t.right+"px",paddingBottom:t.footer+t.insetFooter+t.bottom+"px",paddingLeft:t.left+"px"}}},render:function(t){var e={staticClass:"v-content",style:this.styles,ref:"content"};return t(this.tag,e,[t("div",{staticClass:"v-content__wrap"},this.$slots.default)])}}},"71d9":function(t,e,i){"use strict";i("d263"),i("c5f6"),i("ae8d");var n=i("c6f7"),o=i("b64a"),s=i("6a18"),a=i("b57a");var r={inserted:function(t,e){var i=e.value,n=e.options||{passive:!0},o=e.arg?document.querySelector(e.arg):window;o&&(o.addEventListener("scroll",i,n),t._onScroll={callback:i,options:n,target:o})},unbind:function(t){if(t._onScroll){var e=t._onScroll,i=e.callback,n=e.options;e.target.removeEventListener("scroll",i,n),delete t._onScroll}}},c=i("d9bd"),l=i("58df"),h=Object.assign||function(t){for(var e=1;e<arguments.length;e++){var i=arguments[e];for(var n in i)Object.prototype.hasOwnProperty.call(i,n)&&(t[n]=i[n])}return t};e.a=Object(l.a)(Object(n.a)("top",["clippedLeft","clippedRight","computedHeight","invertedScroll","manualScroll"]),o.a,a.a,s.a).extend({name:"v-toolbar",directives:{Scroll:r},props:{card:Boolean,clippedLeft:Boolean,clippedRight:Boolean,dense:Boolean,extended:Boolean,extensionHeight:{type:[Number,String],validator:function(t){return!isNaN(parseInt(t))}},flat:Boolean,floating:Boolean,height:{type:[Number,String],validator:function(t){return!isNaN(parseInt(t))}},invertedScroll:Boolean,manualScroll:Boolean,prominent:Boolean,scrollOffScreen:Boolean,scrollToolbarOffScreen:Boolean,scrollTarget:String,scrollThreshold:{type:Number,default:300},tabs:Boolean},data:function(){return{activeTimeout:null,currentScroll:0,heights:{mobileLandscape:48,mobile:56,desktop:64,dense:48},isActive:!0,isExtended:!1,isScrollingUp:!1,previousScroll:0,savedScroll:0,target:null}},computed:{canScroll:function(){return this.scrollToolbarOffScreen?(Object(c.d)("scrollToolbarOffScreen","scrollOffScreen",this),!0):this.scrollOffScreen||this.invertedScroll},computedContentHeight:function(){return this.height?parseInt(this.height):this.dense?this.heights.dense:this.prominent||this.$vuetify.breakpoint.mdAndUp?this.heights.desktop:this.$vuetify.breakpoint.smAndDown&&this.$vuetify.breakpoint.width>this.$vuetify.breakpoint.height?this.heights.mobileLandscape:this.heights.mobile},computedExtensionHeight:function(){return this.tabs?48:this.extensionHeight?parseInt(this.extensionHeight):this.computedContentHeight},computedHeight:function(){return this.isExtended?this.computedContentHeight+this.computedExtensionHeight:this.computedContentHeight},computedMarginTop:function(){return this.app?this.$vuetify.application.bar:0},classes:function(){return h({"v-toolbar":!0,"elevation-0":this.flat||!this.isActive&&!this.tabs&&this.canScroll,"v-toolbar--absolute":this.absolute,"v-toolbar--card":this.card,"v-toolbar--clipped":this.clippedLeft||this.clippedRight,"v-toolbar--dense":this.dense,"v-toolbar--extended":this.isExtended,"v-toolbar--fixed":!this.absolute&&(this.app||this.fixed),"v-toolbar--floating":this.floating,"v-toolbar--prominent":this.prominent},this.themeClasses)},computedPaddingLeft:function(){return!this.app||this.clippedLeft?0:this.$vuetify.application.left},computedPaddingRight:function(){return!this.app||this.clippedRight?0:this.$vuetify.application.right},computedTransform:function(){return this.isActive?0:this.canScroll?-this.computedContentHeight:-this.computedHeight},currentThreshold:function(){return Math.abs(this.currentScroll-this.savedScroll)},styles:function(){return{marginTop:this.computedMarginTop+"px",paddingRight:this.computedPaddingRight+"px",paddingLeft:this.computedPaddingLeft+"px",transform:"translateY("+this.computedTransform+"px)"}}},watch:{currentThreshold:function(t){this.invertedScroll?this.isActive=this.currentScroll>this.scrollThreshold:t<this.scrollThreshold||!this.isBooted||(this.isActive=this.isScrollingUp,this.savedScroll=this.currentScroll)},isActive:function(){this.savedScroll=0},invertedScroll:function(t){this.isActive=!t},manualScroll:function(t){this.isActive=!t},isScrollingUp:function(){this.savedScroll=this.savedScroll||this.currentScroll}},created:function(){(this.invertedScroll||this.manualScroll)&&(this.isActive=!1)},mounted:function(){this.scrollTarget&&(this.target=document.querySelector(this.scrollTarget))},methods:{onScroll:function(){this.canScroll&&!this.manualScroll&&"undefined"!=typeof window&&(this.currentScroll=this.target?this.target.scrollTop:window.pageYOffset,this.isScrollingUp=this.currentScroll<this.previousScroll,this.previousScroll=this.currentScroll)},updateApplication:function(){return this.invertedScroll||this.manualScroll?0:this.computedHeight}},render:function(t){this.isExtended=this.extended||!!this.$slots.extension;var e=[],i=this.setBackgroundColor(this.color,{class:this.classes,style:this.styles,on:this.$listeners});return i.directives=[{arg:this.scrollTarget,name:"scroll",value:this.onScroll}],e.push(t("div",{staticClass:"v-toolbar__content",style:{height:this.computedContentHeight+"px"},ref:"content"},this.$slots.default)),this.isExtended&&e.push(t("div",{staticClass:"v-toolbar__extension",style:{height:this.computedExtensionHeight+"px"}},this.$slots.extension)),t("nav",i,e)}})},"98a8":function(t,e,i){"use strict";i.r(e);var n,o=i("d225"),s=i("b0b4"),a=i("308d"),r=i("6bb5"),c=i("4e2b"),l=i("60a3"),h=Object(l.a)({name:"SimpleLayout"})(n=function(t){function e(){return Object(o.a)(this,e),Object(a.a)(this,Object(r.a)(e).apply(this,arguments))}return Object(c.a)(e,t),Object(s.a)(e,[{key:"created",value:function(){console.log("Simple-Layout")}}]),e}(l.b))||n,u=i("2877"),d=i("6544"),p=i.n(d),f=i("a523"),v=i("549c"),b=Object(u.a)(h,function(){var t=this.$createElement,e=this._self._c||t;return e("v-content",[e("v-container",{attrs:{fluid:""}},[e("router-view")],1)],1)},[],!1,null,null,null);e.default=b.exports,p()(b,{VContainer:f.a,VContent:v.a})},a523:function(t,e,i){"use strict";i("db6d");var n=i("e8f2");e.a=Object(n.a)("container")},ae8d:function(t,e,i){},b57a:function(t,e,i){"use strict";var n=i("2b0e");e.a=n.default.extend({name:"ssr-bootable",data:function(){return{isBooted:!1}},mounted:function(){var t=this;window.requestAnimationFrame(function(){t.$el.setAttribute("data-booted","true"),t.isBooted=!0})}})},c6f7:function(t,e,i){"use strict";i.d(e,"a",function(){return s});var n=i("c22b"),o=i("58df");function s(t){var e=arguments.length>1&&void 0!==arguments[1]?arguments[1]:[];return Object(o.a)(Object(n.b)(["absolute","fixed"])).extend({name:"applicationable",props:{app:Boolean},computed:{applicationProperty:function(){return t}},watch:{app:function(t,e){e?this.removeApplication(!0):this.callUpdate()},applicationProperty:function(t,e){this.$vuetify.application.unbind(this._uid,e)}},activated:function(){this.callUpdate()},created:function(){for(var t=0,i=e.length;t<i;t++)this.$watch(e[t],this.callUpdate);this.callUpdate()},mounted:function(){this.callUpdate()},deactivated:function(){this.removeApplication()},destroyed:function(){this.removeApplication()},methods:{callUpdate:function(){this.app&&this.$vuetify.application.bind(this._uid,this.applicationProperty,this.updateApplication())},removeApplication:function(){(arguments.length>0&&void 0!==arguments[0]&&arguments[0]||this.app)&&this.$vuetify.application.unbind(this._uid,this.applicationProperty)},updateApplication:function(){return 0}}})}},dead:function(t,e,i){"use strict";i.r(e);var n,o=i("d225"),s=i("b0b4"),a=i("308d"),r=i("6bb5"),c=i("4e2b"),l=i("60a3"),h=Object(l.a)({name:"AppLayout",components:{Sidebar:function(){return i.e("sidebar").then(i.bind(null,"5ea5"))}}})(n=function(t){function e(){return Object(o.a)(this,e),Object(a.a)(this,Object(r.a)(e).apply(this,arguments))}return Object(c.a)(e,t),Object(s.a)(e,[{key:"created",value:function(){console.log("App-Layout")}}]),e}(l.b))||n,u=i("2877"),d=i("6544"),p=i.n(d),f=i("a523"),v=i("549c"),b=(i("d263"),i("c5f6"),i("e93b"),i("c6f7")),g=i("b64a"),m=i("6a18"),S=Object.assign||function(t){for(var e=1;e<arguments.length;e++){var i=arguments[e];for(var n in i)Object.prototype.hasOwnProperty.call(i,n)&&(t[n]=i[n])}return t},y={name:"v-footer",mixins:[Object(b.a)(null,["height","inset"]),g.a,m.a],props:{height:{default:32,type:[Number,String]},inset:Boolean},computed:{applicationProperty:function(){return this.inset?"insetFooter":"footer"},computedMarginBottom:function(){if(this.app)return this.$vuetify.application.bottom},computedPaddingLeft:function(){return this.app&&this.inset?this.$vuetify.application.left:0},computedPaddingRight:function(){return this.app&&this.inset?this.$vuetify.application.right:0},styles:function(){var t={height:isNaN(this.height)?this.height:this.height+"px"};return this.computedPaddingLeft&&(t.paddingLeft=this.computedPaddingLeft+"px"),this.computedPaddingRight&&(t.paddingRight=this.computedPaddingRight+"px"),this.computedMarginBottom&&(t.marginBottom=this.computedMarginBottom+"px"),t}},methods:{updateApplication:function(){var t=parseInt(this.height);return isNaN(t)?this.$el?this.$el.clientHeight:0:t}},render:function(t){return t("footer",this.setBackgroundColor(this.color,{staticClass:"v-footer",class:S({"v-footer--absolute":this.absolute,"v-footer--fixed":!this.absolute&&(this.app||this.fixed),"v-footer--inset":this.inset},this.themeClasses),style:this.styles,ref:"content"}),this.$slots.default)}},x=i("71d9"),O=Object(u.a)(h,function(){var t=this.$createElement,e=this._self._c||t;return e("div",[e("v-toolbar",{attrs:{app:""}}),e("Sidebar"),e("v-content",[e("v-container",{attrs:{fluid:""}},[e("router-view")],1)],1),e("v-footer",{attrs:{app:""}})],1)},[],!1,null,null,null);e.default=O.exports,p()(O,{VContainer:f.a,VContent:v.a,VFooter:y,VToolbar:x.a})},e93b:function(t,e,i){},f134:function(t,e,i){}}]);