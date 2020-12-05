import * as React from 'react';
import renderer from 'react-test-renderer';
import Login from './Login';
import {
  BrowserRouter as Router,
} from "react-router-dom";

it("login snapshot", () => {
  const tree = renderer.create(
    <Router>
      <Login />
    </Router>
  ).toJSON()
  expect(tree).toMatchSnapshot()
})
