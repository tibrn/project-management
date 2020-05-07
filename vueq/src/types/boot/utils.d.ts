
interface VM {
  [id: string]: any;
}
export interface RequestUtil {
  vm: VM;
  call: Function;
  loading?: string;
  errors?: string;
  debug?: boolean;
}
export interface UtilsInterface {
  readonly request: (object: RequestUtil) => {};

}
