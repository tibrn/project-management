
export interface RequestUtil { vm: any; loading?: string; errors?: string; call: Function }
export interface UtilsInterface {
  readonly request: (object: RequestUtil) => {};

}
