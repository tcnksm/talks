rrset "api.example.com", "A" do
  set_identifier "Monolith"
  weight 255
  resource_records(
    "x.x.x.x",
      ...
  )
end
  
rrset "api.example.com", "A" do
  set_identifier "Gateway"
  weight 1
  resource_records(
    "y.y.y.y",
      ...
  )
end

