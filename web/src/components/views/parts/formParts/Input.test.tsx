import React from 'react';
import { mount } from 'enzyme';
import renderer from 'react-test-renderer';
import Input from "./Input";
import { render } from '@testing-library/react';

describe("Input Test", () => {
  it("value 確認", () => {
    const testState = ""
    const setTestState: React.Dispatch<React.SetStateAction<string>> = () => { }
    const wrapper = mount(<Input value={testState} type="text" setState={setTestState} />)
    console.log(wrapper.debug())
    const inputWrapper = wrapper.find('input[type="text"]')
    expect(inputWrapper.prop("value")).toBe(testState)
  })
  it("onChangeでsetStateが呼ばれる", () => {
    const testState = ""
    const setTestState: React.Dispatch<React.SetStateAction<string>> = jest.fn()
    const wrapper = mount(<Input value={testState} type="text" setState={setTestState} />)
    const inputWrapper = wrapper.find('input[type="text"]').simulate('change')
    expect(setTestState).toHaveBeenCalled()
  })
})
