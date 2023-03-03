/* eslint-disable */
import type * as Types from '../../../../@types'

export type Methods = {
  get: {
    status: 200
    /** Success to get history */
    resBody: Types.External2[]
  }

  post: {
    status: 200
    /** Success to receive */
    resBody: Types.External3

    reqBody: {
      status: 'before_started' | 'in_progress' | 'success' | 'failed'
    }
  }
}
