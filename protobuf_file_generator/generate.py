import example_pb2  # Import the generated protobuf module

def create_and_serialize_pb():
    # Create a Protobuf message instance
    data = example_pb2.MyMessage()

    # Set the values of the message
    data.name = "John Doe"
    data.age = 30

    # Serialize the Protobuf message to binary format
    serialized_data = data.SerializeToString()

    # Save the binary data to a file
    with open("data.pb", "wb") as f:
        f.write(serialized_data)

if __name__ == "__main__":
    create_and_serialize_pb()