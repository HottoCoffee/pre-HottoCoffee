import type { AspidaClient, BasicHeaders } from 'aspida'
import { dataToURLString } from 'aspida'
import type { Methods as Methods0 } from './api/batch'
import type { Methods as Methods1 } from './api/batch/_batch_id@number'
import type { Methods as Methods2 } from './api/batch/_batch_id@number/history'
import type { Methods as Methods3 } from './api/batch/_id@number/history/_historyId@number'
import type { Methods as Methods4 } from './api/calendar'

const api = <T>({ baseURL, fetch }: AspidaClient<T>) => {
  const prefix = (baseURL === undefined ? 'http://localhost:3000' : baseURL).replace(/\/$/, '')
  const PATH0 = '/api/batch'
  const PATH1 = '/history'
  const PATH2 = '/api/calendar'
  const GET = 'GET'
  const POST = 'POST'
  const PUT = 'PUT'
  const DELETE = 'DELETE'

  return {
    api: {
      batch: {
        _batch_id: (val2: number) => {
          const prefix2 = `${PATH0}/${val2}`

          return {
            history: {
              /**
               * @returns success to get history
               */
              get: (option?: { config?: T | undefined } | undefined) =>
                fetch<Methods2['get']['resBody'], BasicHeaders, Methods2['get']['status']>(prefix, `${prefix2}${PATH1}`, GET, option).json(),
              /**
               * @returns success to get history
               */
              $get: (option?: { config?: T | undefined } | undefined) =>
                fetch<Methods2['get']['resBody'], BasicHeaders, Methods2['get']['status']>(prefix, `${prefix2}${PATH1}`, GET, option).json().then(r => r.body),
              /**
               * @returns success to receive
               */
              post: (option: { body: Methods2['post']['reqBody'], config?: T | undefined }) =>
                fetch<Methods2['post']['resBody'], BasicHeaders, Methods2['post']['status']>(prefix, `${prefix2}${PATH1}`, POST, option).json(),
              /**
               * @returns success to receive
               */
              $post: (option: { body: Methods2['post']['reqBody'], config?: T | undefined }) =>
                fetch<Methods2['post']['resBody'], BasicHeaders, Methods2['post']['status']>(prefix, `${prefix2}${PATH1}`, POST, option).json().then(r => r.body),
              $path: () => `${prefix}${prefix2}${PATH1}`
            },
            /**
             * @returns successfully deleted
             */
            delete: (option?: { config?: T | undefined } | undefined) =>
              fetch<Methods1['delete']['resBody'], BasicHeaders, Methods1['delete']['status']>(prefix, prefix2, DELETE, option).json(),
            /**
             * @returns successfully deleted
             */
            $delete: (option?: { config?: T | undefined } | undefined) =>
              fetch<Methods1['delete']['resBody'], BasicHeaders, Methods1['delete']['status']>(prefix, prefix2, DELETE, option).json().then(r => r.body),
            /**
             * @returns successfully updated.
             */
            put: (option: { body: Methods1['put']['reqBody'], config?: T | undefined }) =>
              fetch<Methods1['put']['resBody'], BasicHeaders, Methods1['put']['status']>(prefix, prefix2, PUT, option).json(),
            /**
             * @returns successfully updated.
             */
            $put: (option: { body: Methods1['put']['reqBody'], config?: T | undefined }) =>
              fetch<Methods1['put']['resBody'], BasicHeaders, Methods1['put']['status']>(prefix, prefix2, PUT, option).json().then(r => r.body),
            /**
             * @returns return specific batch information
             */
            get: (option?: { config?: T | undefined } | undefined) =>
              fetch<Methods1['get']['resBody'], BasicHeaders, Methods1['get']['status']>(prefix, prefix2, GET, option).json(),
            /**
             * @returns return specific batch information
             */
            $get: (option?: { config?: T | undefined } | undefined) =>
              fetch<Methods1['get']['resBody'], BasicHeaders, Methods1['get']['status']>(prefix, prefix2, GET, option).json().then(r => r.body),
            $path: () => `${prefix}${prefix2}`
          }
        },
        _id: (val2: number) => {
          const prefix2 = `${PATH0}/${val2}`

          return {
            history: {
              _historyId: (val4: number) => {
                const prefix4 = `${prefix2}${PATH1}/${val4}`

                return {
                  /**
                   * @returns Success to get history
                   */
                  get: (option?: { config?: T | undefined } | undefined) =>
                    fetch<Methods3['get']['resBody'], BasicHeaders, Methods3['get']['status']>(prefix, prefix4, GET, option).json(),
                  /**
                   * @returns Success to get history
                   */
                  $get: (option?: { config?: T | undefined } | undefined) =>
                    fetch<Methods3['get']['resBody'], BasicHeaders, Methods3['get']['status']>(prefix, prefix4, GET, option).json().then(r => r.body),
                  $path: () => `${prefix}${prefix4}`
                }
              }
            }
          }
        },
        /**
         * get batch list. response can be filtered by keyword. If no batches are found, an empty array is returned.
         * @returns success filter and can get batch list
         */
        get: (option?: { query?: Methods0['get']['query'] | undefined, config?: T | undefined } | undefined) =>
          fetch<Methods0['get']['resBody'], BasicHeaders, Methods0['get']['status']>(prefix, PATH0, GET, option).json(),
        /**
         * get batch list. response can be filtered by keyword. If no batches are found, an empty array is returned.
         * @returns success filter and can get batch list
         */
        $get: (option?: { query?: Methods0['get']['query'] | undefined, config?: T | undefined } | undefined) =>
          fetch<Methods0['get']['resBody'], BasicHeaders, Methods0['get']['status']>(prefix, PATH0, GET, option).json().then(r => r.body),
        /**
         * register new batch
         * @returns success to create
         */
        post: (option: { body: Methods0['post']['reqBody'], config?: T | undefined }) =>
          fetch<Methods0['post']['resBody'], BasicHeaders, Methods0['post']['status']>(prefix, PATH0, POST, option).json(),
        /**
         * register new batch
         * @returns success to create
         */
        $post: (option: { body: Methods0['post']['reqBody'], config?: T | undefined }) =>
          fetch<Methods0['post']['resBody'], BasicHeaders, Methods0['post']['status']>(prefix, PATH0, POST, option).json().then(r => r.body),
        $path: (option?: { method?: 'get' | undefined; query: Methods0['get']['query'] } | undefined) =>
          `${prefix}${PATH0}${option && option.query ? `?${dataToURLString(option.query)}` : ''}`
      },
      calendar: {
        /**
         * @returns return batch history list
         */
        get: (option: { query: Methods4['get']['query'], config?: T | undefined }) =>
          fetch<Methods4['get']['resBody'], BasicHeaders, Methods4['get']['status']>(prefix, PATH2, GET, option).json(),
        /**
         * @returns return batch history list
         */
        $get: (option: { query: Methods4['get']['query'], config?: T | undefined }) =>
          fetch<Methods4['get']['resBody'], BasicHeaders, Methods4['get']['status']>(prefix, PATH2, GET, option).json().then(r => r.body),
        $path: (option?: { method?: 'get' | undefined; query: Methods4['get']['query'] } | undefined) =>
          `${prefix}${PATH2}${option && option.query ? `?${dataToURLString(option.query)}` : ''}`
      }
    }
  }
}

export type ApiInstance = ReturnType<typeof api>
export default api
