import React from "react";
import { connect } from 'react-redux';

import { ObjectView } from '../../shared/objectView';
import style from './result.module.css';

class ResultComponent extends React.Component {
  render() {
    if(Object.keys(this.props.result).length === 0) {
      return (
        <div className={style.preparing}>
          <p>Please upload your file first</p>
        </div>
      )
    }
    return (
      <ObjectView object={this.props.result} property={'Result'}></ObjectView>
    )
  }
}

function mapStateTpProps(state) {
  return {
    result: state
  }
}

export const Result = connect(mapStateTpProps)(ResultComponent)

