(window.jsonpFunction=window.jsonpFunction||[]).push([["projects"],{aaa9:function(e,t,n){"use strict";n.r(t);var a=(n("96cf"),n("3b8d")),r=n("d225"),s=n("b0b4"),c=n("308d"),o=n("6bb5"),i=n("4e2b"),u=n("9ab4"),l=n("60a3"),p=function(e){function t(){var e;return Object(r.a)(this,t),(e=Object(c.a)(this,Object(o.a)(t).apply(this,arguments))).isLoading=!1,e.projects=[],e}return Object(i.a)(t,e),Object(s.a)(t,[{key:"created",value:function(){console.log("CEVA")}},{key:"mounted",value:function(){console.log("GET"),this.getTasks()}},{key:"getTasks",value:function(){var e=Object(a.a)(regeneratorRuntime.mark(function e(){var t,n;return regeneratorRuntime.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return this.isLoading=!0,e.prev=1,e.next=4,this.axios.get("/api/project");case 4:t=e.sent,n=t.data,console.log(n),e.next=12;break;case 9:e.prev=9,e.t0=e.catch(1),console.log("ProjectsIndex",e.t0);case 12:return e.prev=12,this.isLoading=!1,e.finish(12);case 15:case"end":return e.stop()}},e,this,[[1,9,12,15]])}));return function(){return e.apply(this,arguments)}}()}]),t}(l.b),b=p=u.a([Object(l.a)({name:"ProjectsIndex"})],p),d=n("2877"),f=Object(d.a)(b,function(){var e=this.$createElement;return(this._self._c||e)("div")},[],!1,null,null,null);t.default=f.exports}}]);