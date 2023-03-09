/* eslint-disable */
export type Batch = External0_Batch

export type External0_Batch = {
  id: number
  batch_name: string
  server_name: string
  /** ex.) 30 * * * * */
  cron_setting: string
  initial_date: string
  /** Unit -> Minutes */
  time_limit: number
  /** Unit -> Minutes */
  estimated_duration?: number | undefined
}

/** Error Schema */
export type External1_Error = {
  /** Error message */
  message: string
}

/** Success Schema */
export type External2_Success = {
  message: string
}

export type External3_History = {
  history_id?: number | undefined
  batch_id: number
  batch_name: string
  start_datetime: string
  status: 'before_started' | 'in_progress' | 'success' | 'failed'
}
