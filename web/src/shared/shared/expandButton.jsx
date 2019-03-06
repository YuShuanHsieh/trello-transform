import React from 'react';
import PropTypes from 'prop-types';
import { ExpandMore, ExpandLess } from '@material-ui/icons';
import { Button } from '@material-ui/core';

export class ExpandButton extends React.Component {
  static propTypes = {
    onToggle: PropTypes.func,
    expand: PropTypes.bool.isRequired,
    label: PropTypes.string.isRequired,
  }

  static defaultProps = {
    onToggle: () => {},
  }

  handleExpand = () => {
    this.props.onToggle(!this.props.expand);
  }

  render() {
    return (
      <Button onClick={this.handleExpand}>
        {this.props.expand ? <ExpandLess /> : <ExpandMore />}
        {this.props.label}
      </Button>
    );
  }
}
