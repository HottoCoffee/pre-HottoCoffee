/* eslint-disable */
export type Batch = External0

export type External0 = {
  id: number
  batch_name: string
  server_name: string
  initial_date: string
  time_limit: number
  estimated_duration?: number | undefined
  status?: 'before_started' | 'in_progress' | 'success' | 'failed' | undefined
}

/** Error Schema */
export type External1 = {
  /** type of error */
  state: string
  /** Error message */
  message: string
}

export type External2 = {
  id: number
  batch_id: number
  status: 'success' | 'failure'
}

/** Success Schema */
export type External3 = {
  message: string
}
