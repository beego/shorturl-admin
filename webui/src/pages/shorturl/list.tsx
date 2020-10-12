// @BeeOverwrite YES
// @BeeGenerateTime 20201012_230803
import { Card, message, Tag, Button, Divider, Modal } from 'antd';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import request from "@/utils/request";
import ProTable, { ProColumns, ActionType } from '@ant-design/pro-table';
import React, { useState, useRef,Fragment } from 'react';
import ListForm from "./components/ListForm"
import { PlusOutlined } from '@ant-design/icons';

const apiUrl = "/api/adminß/shorturl";

const reqList = (params) => {
   return request(apiUrl+"/list", {
        params,
   });
}


const handleCreate = async (values) => {
  const hide = message.loading('正在添加');
  try {
    const resp = await request(apiUrl+"/create" , {
      method: 'POST',
      data: {
        ...values,
      },
    })
    if (resp.code !== 0) {
      hide();
      message.error('添加失败，错误信息：'+resp.msg);
      return true
    }
    hide();
    message.success('添加成功');
    return true;
  } catch (error) {
    hide();
    message.error('添加失败请重试！');
    return false;
  }
};

const handleUpdate = async (values) => {
  const hide = message.loading('正在更新');
  try {
    const resp = await request(apiUrl+"/update" , {
      method: 'POST',
      data: {
        ...values,
      },
    })
    if (resp.code !== 0) {
      hide();
      message.error('更新失败，错误信息：'+resp.msg);
      return true
    }
    hide();
    message.success('更新成功');
    return true;
  } catch (error) {
    hide();
    message.error('更新失败请重试！');
    return false;
  }
};



const TableList: React.FC<{}> = () => {
  const actionRef = useRef<ActionType>();
  const [createModalVisible, handleCreateModalVisible] = useState<boolean>(false);
  const [updateModalVisible, handleUpdateModalVisible] = useState<boolean>(false);
  const [updateValues, setUpdateValues] = useState({});

    const columns = [
        
        
        {
            title: "id",
            dataIndex: "id",
            key: "id",
        },
        
        
        
        {
            title: "标题",
            dataIndex: "title",
            key: "title",
        },
        
        
        
        {
            title: "汇总id",
            dataIndex: "summaryId",
            key: "summaryId",
        },
        
        
        
        {
            title: "url",
            dataIndex: "url",
            key: "url",
        },
        
        
      {
        title: '操作',
        dataIndex: 'operating',
        key: 'operating',
        valueType:"option",
        render: (value, record) => (
          <Fragment>
            <a
              onClick={() => {
                handleUpdateModalVisible(true);
                setUpdateValues(record);
              }}
            >
              编辑
            </a>
            <Divider type="vertical" />
            <a
              onClick={() => {
                Modal.confirm({
                  title: '确认删除？',
                  okText: '确认',
                  cancelText: '取消',
                  onOk: () => {
                    request(apiUrl+"/delete", {
                       method: 'POST',
                       data: {
                          id: record.id,
                       },
                     }).then((res) => {
                      if (res.code !== 0) {
                        message.error(res.msg);
                        return false;
                      }
                      actionRef.current?.reloadAndRest();
                      return true;
                    });
                  },
                });
              }}
            >
              删除
            </a>
          </Fragment>
        ),
      },
    ];
    return (
      <PageHeaderWrapper>
        <Card>
          <ProTable
            actionRef={actionRef}
            request={(params, sorter, filter) => reqList({ ...params, sorter, filter })}
            columns={columns}
            rowKey={(record) => record.id}
            toolBarRender={action => [
                 <Button type="primary" onClick={() => handleCreateModalVisible(true)}>
                   <PlusOutlined /> 新建
                 </Button>,
            ]}
          />
        </Card>
        <ListForm
          formTitle={"创建"}
          onSubmit={async (value) => {
            const success = handleCreate(value);
            if (success) {
              handleCreateModalVisible(false);
              if (actionRef.current) {
                actionRef.current.reload();
              }
            }
          }}
          onCancel={() => handleCreateModalVisible(false)}
          modalVisible={createModalVisible}
        />
        <ListForm
          formTitle={"编辑"}
          onSubmit={async (value) => {
            const success = await handleUpdate(value);
            if (success) {
              handleUpdateModalVisible(false);
              setUpdateValues({});
              if (actionRef.current) {
                actionRef.current.reload();
              }
            }
          }}
          onCancel={() => {
            setUpdateValues({})
            handleUpdateModalVisible(false)
          }}
          modalVisible={updateModalVisible}
          initialValues={updateValues}
        />
      </PageHeaderWrapper>
    );
}
export default TableList;
