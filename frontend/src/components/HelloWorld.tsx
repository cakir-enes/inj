import React, { useState, useEffect } from "react";
import Modal from "react-modal";

let a = {
  showModal: true
};

const HelloWorld: React.FC = () => {
  // const [showModal, setShowModal] = useState(false)
  const [a, setA] = useState(1);

  useEffect(() => {
    let tok = setInterval(() => setA(x => x + 1), 24 / 60);
    return () => clearInterval(tok);
  });
  return <div>{a}</div>;
};

// class HelloWorld extends React.Component {
//   constructor(props: any, context: any) {
//     super(props, context);
//     this.state = {
//       showModal: false
//     };
//     let x = a;
//     this.handleOpenModal = this.handleOpenModal.bind(this);
//     this.handleCloseModal = this.handleCloseModal.bind(this);
//   }

//   handleOpenModal() {
//     this.setState({ showModal: true });

//     window.backend.basic().then(result =>
//       this.setState({
//         result
//       })
//     );
//   }

//   handleCloseModal() {
//     this.setState({ showModal: false });
//   }

//   render() {
//     let { result }: any = this.state;

//     return (
//       <div className="App">
//         <button onClick={this.handleOpenModal} type="button">
//           Hello
//         </button>
//         <Modal
//           isOpen={this.state.showModal}
//           contentLabel="Minimal Modal Example"
//         >
//           <p>{result}</p>
//           <button onClick={this.handleCloseModal}>Close Modal</button>
//         </Modal>
//       </div>
//     );
//   }
// }

export default HelloWorld;
