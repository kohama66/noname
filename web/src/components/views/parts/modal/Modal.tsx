import React, { FC, useState } from 'react';
import "./Modal.scss";

interface props {
  isModal: boolean
  modelText: string
  handleClick: (type: "yes" | "no") => Promise<void>
  buttonDisabled: boolean
}

const Modal: FC<props> = (props) => {
  return (
    <div id="modal" className={props.isModal ? "on-modal" : ""}>
      <div id="modal-contents">
        <h2>{props.modelText}</h2>
        <i className="fas fa-smile"></i>
        <div>
          <button onClick={() => props.handleClick("yes")} disabled={props.buttonDisabled}>Yes</button>
          <button onClick={() => props.handleClick("no")} disabled={props.buttonDisabled}>No</button>
        </div>
      </div>
    </div>
  )
}

export default Modal;