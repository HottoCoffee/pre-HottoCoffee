/* eslint-disable */
import type * as Types from '../../@types'

export type Methods = {
  /** get batch list. response can be filtered by keyword. If no batches are found, an empty array is returned. */
  get: {
    query?: {
      keyword?: string | undefined
    } | undefined

    status: 200
    /** success filter and can get batch list */
    resBody: Types.External0_Batch[]
  }

  /** register new batch */
  post: {
    status: 200
    /** success to create */
    resBody: Types.External0_Batch
    reqBody: Types.External0_Batch
  }
}
