
export interface RequestUtil {
  vm: any;
  call: Function;
  loading?: string;
  errors?: string;
  debug?: boolean;
}
export interface UtilsInterface {
  readonly request: (object: RequestUtil) => {};

}
