export interface Task {
  id: number;
  task_id?: number;
  project_id: number;
  name: string;
  progress: string;
  created_at: string;
  updated_at: string;
  subtasks?: Array<Task>;
}
