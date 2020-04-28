import { Module } from 'webpack'

interface ModuleStatic extends Module {
  default: any;
}

interface StaticRoot {
  example: StaticExample;

}
