/* eslint-disable no-unused-vars */
import React, { Component, createRef } from 'react';
import PropTypes from 'prop-types';
import Chart from 'chart.js';

export class PieChart extends Component {
  static propTypes = {
    title: PropTypes.string.isRequired,
    labels: PropTypes.arrayOf(PropTypes.string).isRequired,
    data: PropTypes.arrayOf(PropTypes.number).isRequired,
  }

  ctx = createRef();

  componentDidMount() {
    const chartData = {
      type: 'pie',
      data: {
        labels: this.props.labels,
        datasets: [{
          label: this.props.title,
          data: this.props.data,
        }],
      },
    };
    const chart = new Chart(this.ctx, {
      type: 'pie',
      data: this.chartData,
    });
  }

  render() {
    return (
      <canvas id="chart" ref={this.ctx} width="400" height="400" />
    );
  }
}
