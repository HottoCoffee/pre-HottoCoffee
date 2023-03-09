/* eslint-disable */
import type * as Types from '../../@types'

export type Methods = {
  get: {
    query: {
      /** query for filtering (inclusive) */
      start_date: string
      /** query for filtering (inclusive) */
      end_date: string
    }

    status: 200
    /** return batch history list */
    resBody: Types.External3_History[]
  }
}
