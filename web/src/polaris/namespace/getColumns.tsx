import * as React from 'react'
import { DuckCmpProps } from 'saga-duck'
import NamespacePageDuck, { NamespaceItem } from './PageDuck'
import { Text } from 'tea-component'
import { Column } from '../common/ducks/GridPage'
import Action from '../common/duckComponents/grid/Action'

export default ({ duck: { creators } }: DuckCmpProps<NamespacePageDuck>): Column<NamespaceItem>[] => [
  {
    key: 'name',
    header: '名称',
    render: x => <Text>{x.name}</Text>,
  },
  {
    key: 'commnet',
    header: '描述',
    render: x => <Text tooltip={x.comment}>{x.comment || '-'}</Text>,
  },
  {
    key: 'owners',
    header: '负责人',
    render: x => <Text>{x.owners ?? '-'}</Text>,
  },
  {
    key: 'ctime',
    header: '创建时间',
    render: x => <Text>{x.ctime}</Text>,
  },
  {
    key: 'mtime',
    header: '修改时间',
    render: x => <Text>{x.mtime}</Text>,
  },
  {
    key: 'action',
    header: '操作',
    render: x => {
      return (
        <React.Fragment>
          <Action fn={dispatch => dispatch(creators.edit(x))} tip={'编辑'}>
            {'编辑'}
          </Action>
          <Action fn={dispatch => dispatch(creators.remove(x))} tip={'删除'}>
            {'删除'}
          </Action>
        </React.Fragment>
      )
    },
  },
]
