/* eslint-disable */
import type * as Types from '../../@types'

export type Methods = {
  get: {
    query?: {
      /** query for filtering */
      start_date?: string | undefined
      /** query for filtering */
      end_date?: string | undefined
    } | undefined

    status: 200
    /** Return batch history list */
    resBody: Types.External2[]
  }
}
