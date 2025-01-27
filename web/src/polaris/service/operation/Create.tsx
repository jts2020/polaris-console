import React from "react";
import { DuckCmpProps, purify } from "saga-duck";
import Duck from "./CreateDuck";
import { Form, Segment, Select, Text, Icon, Bubble } from "tea-component";
import FormField from "@src/polaris/common/duckComponents/form/Field";
import Input from "@src/polaris/common/duckComponents/form/Input";
import Dialog from "@src/polaris/common/duckComponents/Dialog";
import Switch from "@src/polaris/common/duckComponents/form/Switch";

export default function Create(props: DuckCmpProps<Duck>) {
  const { duck, store, dispatch } = props;
  const { selectors } = duck;
  const visible = selectors.visible(store);
  if (!visible) {
    return <noscript />;
  }
  const data = selectors.data(store);
  return (
    <Dialog
      duck={duck}
      store={store}
      dispatch={dispatch}
      size="l"
      title={data.name ? "编辑服务" : "新建服务"}
    >
      <CreateForm duck={duck} store={store} dispatch={dispatch} />
    </Dialog>
  );
}

const CreateForm = purify(function CreateForm(props: DuckCmpProps<Duck>) {
  const { duck, store, dispatch } = props;
  const {
    ducks: { form },
    selectors,
  } = duck;

  const formApi = form.getAPI(store, dispatch);
  const {
    namespace,
    name,
    comment,
    metadata,
    enableNearby,
    ports,
    department,
    business,
  } = formApi.getFields([
    "namespace",
    "name",
    "comment",
    "metadata",
    "enableNearby",
    "ports",
    "business",
    "department",
  ]);
  const options = selectors.options(store);
  return (
    <>
      <Form>
        <FormField field={namespace} label="命名空间" required>
          <Select
            value={namespace.getValue()}
            options={options.namespaceList}
            onChange={(value) => namespace.setValue(value)}
            type={"simulate"}
            appearance={"button"}
            size="l"
          ></Select>
        </FormField>

        <FormField field={name} label={"服务名"} required>
          <Input
            field={name}
            maxLength={128}
            placeholder={"允许数字、英文字母、.、-、_，限制128个字符"}
            size={"l"}
          />
        </FormField>
        <FormField field={department} label={"部门"}>
          <Input field={department} size={"l"} />
        </FormField>
        <FormField field={business} label={"业务"}>
          <Input field={business} size={"l"} />
        </FormField>
        <FormField
          field={metadata}
          label={
            <>
              <Text>服务标签</Text>
              <Bubble
                content={"服务标签可用于标识服务的用处、特征，格式为key:value"}
              >
                <Icon type={"info"}></Icon>
              </Bubble>
            </>
          }
        >
          <Input
            field={metadata}
            placeholder={
              "每个key最长不超过128个字符，每个value最长不超过4096个字符,标签数量不能超过64个"
            }
            multiline
            size={"l"}
          />
        </FormField>
        <FormField field={comment} label={"描述"}>
          <Input
            field={comment}
            maxLength={1024}
            placeholder={"长度不超过1024个字符,标签数量不能超过64个"}
            size={"l"}
          />
        </FormField>
        <FormField field={enableNearby} label={"开启就近访问"}>
          <Switch field={enableNearby} />
        </FormField>
      </Form>
    </>
  );
});
