/* eslint-disable */
import type * as Types from '../../../../@types'

export type Methods = {
  get: {
    status: 200
    /** success to get history */
    resBody: Types.External3_History[]
  }

  post: {
    status: 200
    /** success to receive */
    resBody: Types.External2_Success

    reqBody: {
      status: 'before_started' | 'in_progress' | 'success' | 'failed'
      message: string
    }
  }
}
