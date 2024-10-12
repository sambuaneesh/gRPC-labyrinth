import grpc
import sys
import os


# Add the path to the generated proto files
current_dir = os.path.dirname(os.path.abspath(__file__))
generated_proto_dir = os.path.join(os.path.dirname(current_dir), "generated", "py")
sys.path.append(generated_proto_dir)

from proto import labyrinth_pb2
from proto import labyrinth_pb2_grpc


def get_labyrinth_info(stub):
    # Create an empty request using the existing GetLabyrinthInfoRequest
    response = stub.GetLabyrinthInfo(labyrinth_pb2.GetLabyrinthInfoRequest())
    print(f"Labyrinth size: {response.width}x{response.height}")


def get_player_status(stub):
    # Create an empty request using the existing GetPlayerStatusRequest
    response = stub.GetPlayerStatus(labyrinth_pb2.GetPlayerStatusRequest())
    print(f"Score: {response.score}")
    print(f"Health: {response.health}")
    print(f"Position: ({response.position.x}, {response.position.y})")
    print(f"Remaining spells: {response.remaining_spells}")


def register_move(stub):
    direction = input("Enter move direction (UP/DOWN/LEFT/RIGHT): ").upper()
    response = stub.RegisterMove(
        labyrinth_pb2.RegisterMoveRequest(
            direction=labyrinth_pb2.Direction.Value(direction)
        )
    )
    print(f"Move result: {response.status}")


def revelio(stub):
    x = int(input("Enter target X coordinate: "))
    y = int(input("Enter target Y coordinate: "))
    tile_type = input("Enter tile type to reveal (EMPTY/WALL/COIN): ").upper()
    request = labyrinth_pb2.RevelioRequest(
        target=labyrinth_pb2.Position(x=x, y=y),
        tile_type=labyrinth_pb2.TileType.Value(tile_type),
    )
    responses = stub.Revelio(request)
    for response in responses:
        print(f"Revealed tile at: ({response.position.x}, {response.position.y})")


def bombarda(stub):
    targets = []
    for _ in range(3):
        x = int(input(f"Enter target {_+1} X coordinate: "))
        y = int(input(f"Enter target {_+1} Y coordinate: "))
        targets.append(labyrinth_pb2.Position(x=x, y=y))

    def generate_requests():
        for target in targets:
            yield labyrinth_pb2.BombardaRequest(target=target)

    responses = stub.Bombarda(generate_requests())
    for response in responses:
        print(
            f"Bombarda at ({response.affected_position.x}, {response.affected_position.y}): Success={response.success}"
        )


def main():
    channel = grpc.insecure_channel("localhost:8080")
    stub = labyrinth_pb2_grpc.LabyrinthGameStub(channel)

    while True:
        print("\n--- Labyrinth Game Client ---")
        print("1. Get Labyrinth Info")
        print("2. Get Player Status")
        print("3. Make a Move")
        print("4. Cast Revelio")
        print("5. Cast Bombarda")
        print("6. Exit")

        choice = input("Enter your choice (1-6): ")

        if choice == "1":
            get_labyrinth_info(stub)
        elif choice == "2":
            get_player_status(stub)
        elif choice == "3":
            register_move(stub)
        elif choice == "4":
            revelio(stub)
        elif choice == "5":
            bombarda(stub)
        elif choice == "6":
            break
        else:
            print("Invalid choice. Please try again.")


if __name__ == "__main__":
    main()
