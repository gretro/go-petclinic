import { Button } from "~/shell/components/button";
import { PageHeader } from "~/shell/components/page-header";

const OwnersPage = () => {
  return (
    <>
      <PageHeader title="Owners & Pets Management" subTitle="Consult and manage the clinic's pet Owners and their Pets">
        <div>
          <Button flavour="primary">Add Owner</Button>
        </div>
      </PageHeader>
    </>
  );
};

export default OwnersPage;
