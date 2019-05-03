import { RawLocation } from "vue-router";
import Vue from "vue";
type Next = (to?: RawLocation | false | ((vm: Vue) => any) | void) => any;

export default Next;
