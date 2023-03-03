/* eslint-disable */
import type * as Types from '../../@types'

export type Methods = {
  /** Get batch list. Response can be filtered by keyword. */
  get: {
    query?: {
      keyword?: string | undefined
    } | undefined

    status: 200
    /** Success filter and can get batch list */
    resBody: Types.External0[]
  }

  /** Register new batch */
  post: {
    status: 200
    /** Success to create */
    resBody: Types.External0
    reqBody: Types.External0
  }
}
