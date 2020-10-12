// @BeeOverwrite YES
// @BeeGenerateTime 20201012_230803
import {Form, Input, Modal} from 'antd';
import React, { useEffect } from "react";
interface ListFormProps {
  modalVisible: boolean;
  formTitle: string;
  initialValues: {};
  onSubmit: () => void;
  onCancel: () => void;
}

const formLayout = {
  labelCol: { span: 7 },
  wrapperCol: { span: 13 },
};

const ListForm: React.FC<ListFormProps> = (props) => {
  const { modalVisible, onCancel, onSubmit, initialValues, formTitle } = props;
  const [form] = Form.useForm();

  useEffect(() => {
    if (form && !modalVisible) {
      form.resetFields();
    }
  }, [modalVisible]);

  useEffect(() => {
    if (initialValues) {
      form.setFieldsValue({
        ...initialValues,
      });
    }
  }, [initialValues]);

  const handleSubmit = () => {
    if (!form) return;
    form.submit();
  };

  const modalFooter = { okText: '保存', onOk: handleSubmit, onCancel:onCancel }

  return (
      <Modal
        destroyOnClose
        title={formTitle}
        visible={modalVisible}
        {...modalFooter}
      >
        <Form
            {...formLayout}
            form={form}
            onFinish={onSubmit}
            scrollToFirstError
        >
            
            
              <Form.Item
                  name="id"
                  label="id"
                >
                  <Input />
                </Form.Item>
            
            
            
              <Form.Item
                  name="title"
                  label="标题"
                >
                  <Input />
                </Form.Item>
            
            
            
              <Form.Item
                  name="summaryId"
                  label="汇总id"
                >
                  <Input />
                </Form.Item>
            
            
            
              <Form.Item
                  name="url"
                  label="url"
                >
                  <Input />
                </Form.Item>
            
            
        </Form>
      </Modal>
  );
};
export default ListForm;
