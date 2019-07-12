#include "demos/applications/hipster_shop/reflection.h"

#include <google/protobuf/descriptor.pb.h>

#include "src/common/testing/testing.h"

namespace demos {
namespace hipster_shop {

using ::google::protobuf::FileDescriptorSet;
using ::pl::testing::proto::EqualsProto;
using ::pl::testing::proto::Partially;
using ::testing::ElementsAre;

TEST(GetFileDescriptorSetTest, HasAllServicesAndMessages) {
  FileDescriptorSet fds = GetFileDescriptorSet();
  EXPECT_THAT(fds.file(), ElementsAre(Partially(EqualsProto(R"proto(
      syntax: "proto3"
      name: "demos/applications/hipster_shop/proto/demo.proto"
      package: "hipstershop"
  )proto"))));
}

}  // namespace hipster_shop
}  // namespace demos
