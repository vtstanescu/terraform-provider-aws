package ds_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/directoryservice"
	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfds "github.com/hashicorp/terraform-provider-aws/internal/service/ds"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

func TestAccDSDirectory_basic(t *testing.T) {
	var ds directoryservice.DirectoryDescription
	resourceName := "aws_directory_service_directory.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	domainName := acctest.RandomDomainName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(t)
			acctest.PreCheckDirectoryService(t)
			acctest.PreCheckDirectoryServiceSimpleDirectory(t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, directoryservice.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckDirectoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDirectoryConfig_basic(rName, domainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServiceDirectoryExists(resourceName, &ds),
					resource.TestCheckResourceAttrSet(resourceName, "security_group_id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"password",
				},
			},
		},
	})
}

func TestAccDSDirectory_disappears(t *testing.T) {
	var ds directoryservice.DirectoryDescription
	resourceName := "aws_directory_service_directory.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	domainName := acctest.RandomDomainName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(t)
			acctest.PreCheckDirectoryService(t)
			acctest.PreCheckDirectoryServiceSimpleDirectory(t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, directoryservice.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckDirectoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDirectoryConfig_basic(rName, domainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServiceDirectoryExists(resourceName, &ds),
					acctest.CheckResourceDisappears(acctest.Provider, tfds.ResourceDirectory(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDSDirectory_tags(t *testing.T) {
	var ds directoryservice.DirectoryDescription
	resourceName := "aws_directory_service_directory.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	domainName := acctest.RandomDomainName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(t)
			acctest.PreCheckDirectoryService(t)
			acctest.PreCheckDirectoryServiceSimpleDirectory(t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, directoryservice.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckDirectoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDirectoryConfig_tags1(rName, domainName, "key1", "value1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServiceDirectoryExists(resourceName, &ds),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"password",
				},
			},
			{
				Config: testAccDirectoryConfig_tags2(rName, domainName, "key1", "value1updated", "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServiceDirectoryExists(resourceName, &ds),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
			{
				Config: testAccDirectoryConfig_tags1(rName, domainName, "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServiceDirectoryExists(resourceName, &ds),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
		},
	})
}

func TestAccDSDirectory_microsoft(t *testing.T) {
	var ds directoryservice.DirectoryDescription
	resourceName := "aws_directory_service_directory.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	domainName := acctest.RandomDomainName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t); acctest.PreCheckDirectoryService(t) },
		ErrorCheck:               acctest.ErrorCheck(t, directoryservice.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckDirectoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDirectoryConfig_microsoft(rName, domainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServiceDirectoryExists(resourceName, &ds),
					resource.TestCheckResourceAttr(resourceName, "edition", directoryservice.DirectoryEditionEnterprise),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"password",
				},
			},
		},
	})
}

func TestAccDSDirectory_microsoftStandard(t *testing.T) {
	var ds directoryservice.DirectoryDescription
	resourceName := "aws_directory_service_directory.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	domainName := acctest.RandomDomainName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t); acctest.PreCheckDirectoryService(t) },
		ErrorCheck:               acctest.ErrorCheck(t, directoryservice.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckDirectoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDirectoryConfig_microsoftStandard(rName, domainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServiceDirectoryExists(resourceName, &ds),
					resource.TestCheckResourceAttr(resourceName, "edition", directoryservice.DirectoryEditionStandard),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"password",
				},
			},
		},
	})
}

func TestAccDSDirectory_connector(t *testing.T) {
	var ds directoryservice.DirectoryDescription
	resourceName := "aws_directory_service_directory.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	domainName := acctest.RandomDomainName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(t)
			acctest.PreCheckDirectoryService(t)
			acctest.PreCheckDirectoryServiceSimpleDirectory(t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, directoryservice.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckDirectoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDirectoryConfig_connector(rName, domainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServiceDirectoryExists(resourceName, &ds),
					resource.TestCheckResourceAttrSet(resourceName, "security_group_id"),
					resource.TestCheckResourceAttrSet(resourceName, "connect_settings.0.connect_ips.#"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"password",
				},
			},
		},
	})
}

func TestAccDSDirectory_withAliasAndSSO(t *testing.T) {
	var ds directoryservice.DirectoryDescription
	alias := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_directory_service_directory.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	domainName := acctest.RandomDomainName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(t)
			acctest.PreCheckDirectoryService(t)
			acctest.PreCheckDirectoryServiceSimpleDirectory(t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, directoryservice.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckDirectoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDirectoryConfig_alias(rName, domainName, alias),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServiceDirectoryExists(resourceName, &ds),
					resource.TestCheckResourceAttr(resourceName, "alias", alias),
					resource.TestCheckResourceAttr(resourceName, "enable_sso", "false"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"password",
				},
			},
			{
				Config: testAccDirectoryConfig_sso(rName, domainName, alias),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServiceDirectoryExists(resourceName, &ds),
					resource.TestCheckResourceAttr(resourceName, "alias", alias),
					resource.TestCheckResourceAttr(resourceName, "enable_sso", "true"),
				),
			},
			{
				Config: testAccDirectoryConfig_ssoModified(rName, domainName, alias),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServiceDirectoryExists(resourceName, &ds),
					resource.TestCheckResourceAttr(resourceName, "alias", alias),
					resource.TestCheckResourceAttr(resourceName, "enable_sso", "false"),
				),
			},
		},
	})
}

func testAccCheckDirectoryDestroy(s *terraform.State) error {
	conn := acctest.Provider.Meta().(*conns.AWSClient).DSConn

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_directory_service_directory" {
			continue
		}

		_, err := tfds.FindDirectoryByID(conn, rs.Primary.ID)

		if tfresource.NotFound(err) {
			continue
		}

		if err != nil {
			return err
		}

		return fmt.Errorf("Directory Service Directory %s still exists", rs.Primary.ID)
	}

	return nil
}

func testAccCheckServiceDirectoryExists(n string, v *directoryservice.DirectoryDescription) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Directory Service Directory ID is set")
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).DSConn

		output, err := tfds.FindDirectoryByID(conn, rs.Primary.ID)

		if err != nil {
			return err
		}

		*v = *output

		return nil
	}
}

func testAccDirectoryConfig_basic(rName, domain string) string {
	return acctest.ConfigCompose(
		acctest.ConfigVPCWithSubnets(rName, 2),
		fmt.Sprintf(`
resource "aws_directory_service_directory" "test" {
  name     = %[1]q
  password = "SuperSecretPassw0rd"
  size     = "Small"

  vpc_settings {
    vpc_id     = aws_vpc.test.id
    subnet_ids = aws_subnet.test[*].id
  }
}
`, domain),
	)
}

func testAccDirectoryConfig_tags1(rName, domain, tagKey1, tagValue1 string) string {
	return acctest.ConfigCompose(
		acctest.ConfigVPCWithSubnets(rName, 2),
		fmt.Sprintf(`
resource "aws_directory_service_directory" "test" {
  name     = %[1]q
  password = "SuperSecretPassw0rd"
  size     = "Small"

  vpc_settings {
    vpc_id     = aws_vpc.test.id
    subnet_ids = aws_subnet.test[*].id
  }

  tags = {
    %[2]q = %[3]q
  }
}
`, domain, tagKey1, tagValue1))
}

func testAccDirectoryConfig_tags2(rName, domain, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return acctest.ConfigCompose(
		acctest.ConfigVPCWithSubnets(rName, 2),
		fmt.Sprintf(`
resource "aws_directory_service_directory" "test" {
  name     = %[1]q
  password = "SuperSecretPassw0rd"
  size     = "Small"

  vpc_settings {
    vpc_id     = aws_vpc.test.id
    subnet_ids = aws_subnet.test[*].id
  }

  tags = {
    %[2]q = %[3]q
    %[4]q = %[5]q
  }
}
`, domain, tagKey1, tagValue1, tagKey2, tagValue2))
}

func testAccDirectoryConfig_connector(rName, domain string) string {
	return acctest.ConfigCompose(
		acctest.ConfigVPCWithSubnets(rName, 2),
		fmt.Sprintf(`
resource "aws_directory_service_directory" "test" {
  name     = %[1]q
  password = "SuperSecretPassw0rd"
  size     = "Small"
  type     = "ADConnector"

  connect_settings {
    customer_dns_ips  = aws_directory_service_directory.base.dns_ip_addresses
    customer_username = "Administrator"
    vpc_id            = aws_vpc.test.id
    subnet_ids        = aws_subnet.test[*].id
  }
}

resource "aws_directory_service_directory" "base" {
  name     = %[1]q
  password = "SuperSecretPassw0rd"
  size     = "Small"

  vpc_settings {
    vpc_id     = aws_vpc.test.id
    subnet_ids = aws_subnet.test[*].id
  }
}
`, domain),
	)
}

func testAccDirectoryConfig_microsoft(rName, domain string) string {
	return acctest.ConfigCompose(
		acctest.ConfigVPCWithSubnets(rName, 2),
		fmt.Sprintf(`
resource "aws_directory_service_directory" "test" {
  name     = %[1]q
  password = "SuperSecretPassw0rd"
  type     = "MicrosoftAD"

  vpc_settings {
    vpc_id     = aws_vpc.test.id
    subnet_ids = aws_subnet.test[*].id
  }
}
`, domain),
	)
}

func testAccDirectoryConfig_microsoftStandard(rName, domain string) string {
	return acctest.ConfigCompose(
		acctest.ConfigVPCWithSubnets(rName, 2),
		fmt.Sprintf(`
resource "aws_directory_service_directory" "test" {
  name     = %[1]q
  password = "SuperSecretPassw0rd"
  type     = "MicrosoftAD"
  edition  = "Standard"

  vpc_settings {
    vpc_id     = aws_vpc.test.id
    subnet_ids = aws_subnet.test[*].id
  }
}
`, domain),
	)
}

func testAccDirectoryConfig_alias(rName, domain, alias string) string {
	return acctest.ConfigCompose(
		acctest.ConfigVPCWithSubnets(rName, 2),
		fmt.Sprintf(`
resource "aws_directory_service_directory" "test" {
  name     = %[1]q
  password = "SuperSecretPassw0rd"
  size     = "Small"
  alias    = %[2]q

  vpc_settings {
    vpc_id     = aws_vpc.test.id
    subnet_ids = aws_subnet.test[*].id
  }
}
`, domain, alias),
	)
}

func testAccDirectoryConfig_sso(rName, domain, alias string) string {
	return acctest.ConfigCompose(
		acctest.ConfigVPCWithSubnets(rName, 2),
		fmt.Sprintf(`
resource "aws_directory_service_directory" "test" {
  name       = %[1]q
  password   = "SuperSecretPassw0rd"
  size       = "Small"
  alias      = %[2]q
  enable_sso = true

  vpc_settings {
    vpc_id     = aws_vpc.test.id
    subnet_ids = aws_subnet.test[*].id
  }
}
`, domain, alias),
	)
}

func testAccDirectoryConfig_ssoModified(rName, domain, alias string) string {
	return acctest.ConfigCompose(
		acctest.ConfigVPCWithSubnets(rName, 2),
		fmt.Sprintf(`
resource "aws_directory_service_directory" "test" {
  name       = %[1]q
  password   = "SuperSecretPassw0rd"
  size       = "Small"
  alias      = %[2]q
  enable_sso = false

  vpc_settings {
    vpc_id     = aws_vpc.test.id
    subnet_ids = aws_subnet.test[*].id
  }
}
`, domain, alias),
	)
}
