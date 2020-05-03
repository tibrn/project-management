export interface Project {
  id: number;
  platform_id?: number;
  name: string;
  description: string;
  created_at: string;
  updated_at: string;
  tasks?: Array<object>;
  languages?: Array<object>;
  license?: object;
}
