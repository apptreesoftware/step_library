using System;
using System.Collections.Generic;
using System.Text;

namespace Famis.Model
{
    public class ChildWorkOrder
    {
        public string ExternalId { get; set; }
        public int RequestTypeId { get; set; }
        public int RequestSubTypeId { get; set; }
        public int AssignedToId { get; set; }
        public string StatementOfWork { get; set; }
        public int SpaceId { get; set; }
        public string ParentWOId { get; set; }
        public string RequestorCompanyName { get; set; }
        public string RequestorFirstName { get; set; }
        public string RequestorLastName { get; set; }
        public string RequestorEmail { get; set; }
        public string RequestorPhone { get; set; }
        public int RequestPriorityId { get; set; }
        public int CrewId { get; set; }
        public string ExternalCostCenter { get; set; }
    }
}
