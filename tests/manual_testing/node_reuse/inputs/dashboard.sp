dashboard "inputs" {
  title = "Inputs"

  input "i1" {
    sql = <<-EOQ
          select arn as label, arn as value from aws_account
        EOQ
  }
  input "i2" {
    sql = <<-EOQ
          select arn as label, arn as value from aws_account
        EOQ
  }

  table {
    param "foo" {}
    sql = "select $1"
    args  = {
      arn = self.input.i1.value
    }
  }
  table {
    query = query.q1
    args  = {
      arn = self.input.i2.value
    }
  }

}

query "q1"{
  sql = "select arn from aws_account where arn = $1"
  param "arn" {   }
  search_path="test"
}